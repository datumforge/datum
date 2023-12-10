package marionette

import "go.uber.org/zap"

// Config configures the marionette task manager so that different processes can utilize
// different asynchronous task processing resources depending on process compute constraints
type Config struct {
	Logger     *zap.SugaredLogger
	Workers    int    `default:"4" desc:"the number of workers to process tasks asynchronously"`
	QueueSize  int    `split_words:"true" default:"64" desc:"the number of async tasks to buffer in the queue before blocking"`
	ServerName string `split_words:"true" default:"marionette" desc:"used to describe the marionette service in the log"`
}

// Validate validates the Config instance
func (c Config) Validate() error {
	if c.Workers == 0 {
		return ErrNoWorkers
	}

	if c.ServerName == "" {
		return ErrNoServerName
	}

	return nil
}

// IsZero checks if all the fields of the `Config` instance are set to their zero
// values. If all the fields are zero, it returns `true`, indicating that the `Config`
// instance is considered empty or uninitialized
func (c Config) IsZero() bool {
	return c.Workers == 0 && c.QueueSize == 0 && c.ServerName == ""
}
