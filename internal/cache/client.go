package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config for the redis client
type Config struct {
	// Enabled to enable redis client in the server
	Enabled bool `split_words:"true" default:"true"` // $DATUM_REDIS_ENABLED
	// Address is the host:port to connect to redis
	Address string `split_words:"true" default:"localhost:6379"` // $DATUM_REDIS_ADDRESS
	// Name of the connecting client
	Name string `split_words:"true" default:"datum"` // $DATUM_REDIS_NAME
	// Username to connect to redis
	Username string `split_words:"true" default:""` // $DATUM_REDIS_USERNAME
	// Password, must match the password specified in the server configuration
	Password string `split_words:"true" default:""` // $DATUM_REDIS_PASSWORD
	// DB to be selected after connecting to the server, 0 uses the default
	DB int `split_words:"true" default:"0"` // $DATUM_REDIS_DB
	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `split_words:"true" default:"5s"` // $DATUM_REDIS_DIAL_TIMEOUT
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetReadDeadline calls completely.
	ReadTimeout time.Duration `split_words:"true" default:"0"` // $DATUM_REDIS_READ_TIMEOUT
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.  Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetWriteDeadline calls completely.
	WriteTimeout time.Duration `split_words:"true" default:"0"` // $DATUM_REDIS_WRITE_TIMEOUT
	// MaxRetries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	MaxRetries int `split_words:"true" default:"3"` // $DATUM_REDIS_MAX_RETRIES
	// MinIdleConns is useful when establishing new connection is slow.
	// Default is 0. the idle connections are not closed by default.
	MinIdleConns int `split_words:"true" default:"0"` // $DATUM_REDIS_MIN_IDLE_CONNS
	// Maximum number of idle connections.
	// Default is 0. the idle connections are not closed by default.
	MaxIdleConns int `split_words:"true" default:"0"` // $DATUM_REDIS_MAX_IDLE_CONNS
	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActiveConns int `split_words:"true" default:"0"` // $DATUM_REDIS_MAX_ACTIVE_CONNS
}

// New returns a new redis client based on the configuration settings
func New(c Config) *redis.Client {
	opts := &redis.Options{
		Addr:           c.Address,
		ClientName:     c.Name,
		DB:             c.DB,
		DialTimeout:    c.DialTimeout,
		ReadTimeout:    c.ReadTimeout,
		WriteTimeout:   c.WriteTimeout,
		MaxRetries:     c.MaxRetries,
		MinIdleConns:   c.MinIdleConns,
		MaxIdleConns:   c.MaxIdleConns,
		MaxActiveConns: c.MaxActiveConns,
	}

	if c.Username != "" {
		opts.Username = c.Username
	}

	if c.Password != "" {
		opts.Password = c.Password
	}

	return redis.NewClient(opts)
}

// Healthcheck pings the client to check if the connection is working
func Healthcheck(c *redis.Client) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		// check if its alive
		if err := c.Ping(ctx).Err(); err != nil {
			return err
		}

		return nil
	}
}
