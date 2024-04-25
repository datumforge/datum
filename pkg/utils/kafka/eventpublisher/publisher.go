package eventpublisher

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"

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

	// loggerConfig := logger.Config{
	//	App: "Datum",
	//}
	// wmlogger := logger.NewZapLogger(loggerConfig)
	// tmlogger := logger.NewLogger(wmlogger)

	p.PartitionKey = func(topic string, msg *message.Message) (string, error) { return p.Key, nil }
	kafkaConfig := kafka.PublisherConfig{
		Brokers:               p.Config.Kafka.Addresses,
		Marshaler:             kafka.NewWithPartitioningMarshaler(p.PartitionKey),
		OverwriteSaramaConfig: defaultConfig,
	}

	//	var wmlogger watermill.LoggerAdapter

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
