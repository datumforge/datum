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
	Config       config.Config
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

func (p *EventPublisher) BuildPublisher(conf config.Config, logger watermill.LoggerAdapter) (*kafka.Publisher, error) {
	return buildPublisher(conf, logger)
}

func BuildPublisher(conf config.Config, logger watermill.LoggerAdapter) (*kafka.Publisher, error) {
	return buildPublisher(conf, logger)
}

func (p *EventPublisher) Publish(topic string, msg *message.Message) error {
	return p.Publisher.Publish(topic, msg)
}

func buildPublisher(conf config.Config, logger watermill.LoggerAdapter) (*kafka.Publisher, error) {
	defaultConfig := kafka.DefaultSaramaSyncPublisherConfig()
	defaultConfig.ClientID = "kafkid-producer"

	partitionKey := func(topic string, msg *message.Message) (string, error) { return "key", nil }

	kafkaConfig := kafka.PublisherConfig{
		Brokers:               conf.Kafka.Addresses,
		Marshaler:             kafka.NewWithPartitioningMarshaler(partitionKey),
		OverwriteSaramaConfig: defaultConfig,
	}

	return kafka.NewPublisher(kafkaConfig, logger)
}

func (p *EventPublisher) BuildSubscriber(conf config.Config, logger watermill.LoggerAdapter) (*kafka.Subscriber, error) {
	return buildSubscriber(conf, logger)
}

func BuildSubscriber(conf config.Config, logger watermill.LoggerAdapter) (*kafka.Subscriber, error) {
	return buildSubscriber(conf, logger)
}

func buildSubscriber(conf config.Config, logger watermill.LoggerAdapter) (*kafka.Subscriber, error) {
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

func (p *EventPublisher) Shutdown() error { return nil }

func (p *EventPublisher) SetKey(key string) {
	p.Key = key
}

var localAddress = "proxy:10000"

func (p *EventPublisher) BuildLogger(conf config.Config) (*zap.Logger, error) {
	return buildLogger(conf)
}

func BuildLogger(conf config.Config) (*zap.Logger, error) {
	return buildLogger(conf)
}

func buildLogger(conf config.Config) (*zap.Logger, error) {
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
