package config

type Configuration struct {
	App      string
	AppVer   string
	Env      string
	HTTP     HTTPConfig
	Kafka    KafkaConfig
	Consumer ConsumerConfig
}

type HTTPConfig struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

type KafkaConfig struct {
	Addresses []string
}

type ConsumerConfig struct {
	GroupID          string
	Topics           []string
	OffsetFromNewest bool
	Output           struct {
		Stdout       bool
		FileLocation string
	}
}

type Config struct {
	Address string `yaml:"address" validate:"required"`
	Debug   bool
	Kafka   struct {
		Addresses []string `yaml:"addresses" validate:"required"`
	}
	Consumer struct {
		Enabled bool     `yaml:"enabled"`
		GroupID string   `yaml:"groupId" validate:"required"`
		Topics  []string `yaml:"topics" validate:"required"`
		Output  struct {
			Stdout bool   `yaml:"stdout"`
			File   string `yaml:"file"`
		} `yaml:"output"`
	}
}
