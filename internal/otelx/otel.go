package otelx

import (
	"context"
	"net/url"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
)

const (
	StdOutProvider   = "stdout"
	OTLPHTTPProvider = "otlphttp"
	OTLPGRPCProvider = "otlpgrpc"
)

type Config struct {
	Enabled     bool   `yaml:"enabled" split_words:"true" default:"false"`
	Provider    string `yaml:"provider" split_words:"true" default:"stdout"`
	Environment string `yaml:"environment" split_words:"true" default:"development"`
	StdOut      StdOut
	OTLP        OTLP
}

type StdOut struct {
	Pretty           bool `yaml:"pretty" split_words:"true" default:"true"`
	DisableTimestamp bool `yaml:"disable_timestamp" split_words:"true" default:"false"`
}

type OTLP struct {
	Endpoint    string        `yaml:"endpoint" split_words:"true" default:"localhost:4317"`
	Insecure    bool          `yaml:"insecure" split_words:"true" default:"true"`
	Certificate string        `yaml:"certificate" split_words:"true" default:""`
	Headers     []string      `yaml:"headers" split_words:"true" default:""`
	Compression string        `yaml:"compression" split_words:"true" default:""`
	Timeout     time.Duration `yaml:"timeout" split_words:"true" default:"10s"`
}

func NewTracer(c Config, name string, logger *zap.SugaredLogger) error {
	if !c.Enabled {
		logger.Debug("Tracing disabled")
		return nil
	}

	exp, err := newTraceExporter(c)
	if err != nil {
		logger.Debugw("Failed to create trace exporter", "error", err)
		return err
	}

	opts := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		// Record information about this application in a resource.
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
			attribute.String("environment", c.Environment),
		)),
	}

	if exp != nil {
		opts = append(opts, sdktrace.WithBatcher(exp))
	}

	tp := sdktrace.NewTracerProvider(opts...)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return nil
}

func newTraceExporter(c Config) (sdktrace.SpanExporter, error) {
	switch c.Provider {
	case StdOutProvider:
		return newStdoutProvider(c)
	case OTLPHTTPProvider:
		return newOTLPHTTPProvider(c)
	case OTLPGRPCProvider:
		return newOTLPGRPCProvider(c)
	default:
		return nil, newUnknownProviderError(c.Provider)
	}
}

func newStdoutProvider(c Config) (sdktrace.SpanExporter, error) {
	opts := []stdouttrace.Option{}

	if c.StdOut.Pretty {
		opts = append(opts, stdouttrace.WithPrettyPrint())
	}

	if c.StdOut.DisableTimestamp {
		opts = append(opts, stdouttrace.WithoutTimestamps())
	}

	return stdouttrace.New(opts...)
}

func newOTLPHTTPProvider(c Config) (sdktrace.SpanExporter, error) {
	_, err := url.Parse(c.OTLP.Endpoint)
	if err != nil {
		return nil, newTraceConfigError(err)
	}

	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(c.OTLP.Endpoint),
		otlptracehttp.WithTimeout(c.OTLP.Timeout),
	}

	if c.OTLP.Insecure {
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	return otlptrace.New(context.Background(), otlptracehttp.NewClient(opts...))
}

func newOTLPGRPCProvider(c Config) (sdktrace.SpanExporter, error) {
	_, err := url.Parse(c.OTLP.Endpoint)
	if err != nil {
		return nil, newTraceConfigError(err)
	}

	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(c.OTLP.Endpoint),
		otlptracegrpc.WithTimeout(c.OTLP.Timeout),
	}

	if c.OTLP.Insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	return otlptrace.New(context.Background(), otlptracegrpc.NewClient(opts...))
}
