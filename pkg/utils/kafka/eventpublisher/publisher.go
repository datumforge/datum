package eventpublisher

import (
	"github.com/IBM/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/datumforge/datum/pkg/utils/kafka/config"
)

type EventPublisher struct {
	*kafka.Publisher
	PartitionKey kafka.GeneratePartitionKey
	Key          string
	Config       *config.Configuration
}

func (p *EventPublisher) Startup() error {
	defaultConfig := kafka.DefaultSaramaSyncPublisherConfig()
	defaultConfig.ClientID = "kafkid-producer"

	p.PartitionKey = func(topic string, msg *message.Message) (string, error) { return p.Key, nil }
	kafkaConfig := kafka.PublisherConfig{
		Brokers:               p.Config.Kafka.Addresses,
		Marshaler:             kafka.NewWithPartitioningMarshaler(p.PartitionKey),
		OverwriteSaramaConfig: defaultConfig,
	}

	pub, err := kafka.NewPublisher(kafkaConfig, nil)
	if err != nil {
		return err
	}

	p.Publisher = pub
	p.PartitionKey = func(topic string, msg *message.Message) (string, error) { return p.Key, nil }

	return nil
}

func (p *EventPublisher) Shutdown() error { return nil }

func (p *EventPublisher) SetKey(key string) {
	p.Key = key
}

type Config struct {
	Address string `json:"address" default:"localhost:9092"`
	Debug   bool
	Kafka   struct {
		Addresses []string `json:"addresses"`
	}
	Consumer struct {
		Enabled bool     `json:"enabled"`
		GroupID string   `json:"groupId" default:"test-group"`
		Topics  []string `json:"topics"`
		Output  struct {
			Stdout bool   `json:"stdout"`
			File   string `json:"file"`
		} `json:"output"`
	}
}

var localAddress = "localhost:10000"

func BuildLogger(conf *Config) (*zap.Logger, error) {
	outputPaths := make([]string, 0)
	out := conf.Consumer.Output

	if out.Stdout {
		outputPaths = append(outputPaths, "stdout")
	}

	if out.File != "" {
		outputPaths = append(outputPaths, out.File)
	}

	return zap.Config{
		Encoding:      "json",
		Level:         zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:   outputPaths,
		DisableCaller: true,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}.Build()
}

func BuildPublisher(conf *Config, logger watermill.LoggerAdapter) (*kafka.Publisher, error) {
	var brokers []string
	brokers = append(brokers, localAddress)
	kafkaConf := kafka.PublisherConfig{
		Brokers:   brokers,
		Marshaler: kafka.DefaultMarshaler{},
	}

	return kafka.NewPublisher(kafkaConf, logger)
}

func BuildSubscriber(conf *Config, logger watermill.LoggerAdapter) (*kafka.Subscriber, error) {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	var brokers []string
	brokers = append(brokers, localAddress)
	return kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         conf.Consumer.GroupID,
		},
		logger,
	)
}
