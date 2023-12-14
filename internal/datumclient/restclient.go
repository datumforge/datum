package datumclient

// const (
// 	timeout = 10 * time.Second
// )

// // HTTPDoer implements the standard http.Client interface.
// type HTTPDoer interface {
// 	Do(req *http.Request) (*http.Response, error)
// }

// // RestClient is a datum REST API client
// type RestClient struct {
// 	url        string
// 	logger     *zap.Logger
// 	httpClient HTTPDoer
// }

// // URL returns the datum url
// func (c *RestClient) URL() string {
// 	return c.url
// }

// // Option is a functional configuration option
// type Option func(r *RestClient)

// // WithURL sets the datum API URL
// func WithURL(u string) Option {
// 	return func(r *RestClient) {
// 		r.url = u
// 	}
// }

// // WithLogger sets logger
// func WithLogger(l *zap.Logger) Option {
// 	return func(r *RestClient) {
// 		r.logger = l
// 	}
// }

// // WithHTTPClient overrides the default http client
// func WithHTTPClient(c HTTPDoer) Option {
// 	return func(r *RestClient) {
// 		r.httpClient = c
// 	}
// }

// // NewRestClient returns a new datum client for REST endpoint
// func NewRestClient(opts ...Option) (*RestClient, error) {
// 	client := RestClient{
// 		logger: zap.NewNop(),
// 		httpClient: &http.Client{
// 			Timeout: timeout,
// 		},
// 	}

// 	for _, opt := range opts {
// 		opt(&client)
// 	}

// 	return &client, nil
// }
