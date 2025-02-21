Integrate with otel
```go
package logz

import (
	"context"
	"log/slog"
	"os/user"
	"sync"
	"time"

	"github.com/anzx/xtest/internal/utils/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

const (
	otelEndpoint    = "internal.telemetry.gcpnp.anz:443"
	initialInterval = 2 * time.Second
	maxInterval     = 5 * time.Second
	maxElapsedTime  = 10 * time.Second
	shutdownTimeout = 2 * time.Second
)

var (
	LocalVersion string
	once         sync.Once
)

func init() {
	once.Do(func() {
		var err error
		LocalVersion, err = version.Local()
		if err != nil {
			LocalVersion = "unknown"
		}
	})
}

func newResource() *resource.Resource {
	usr, err := user.Current()
	if err != nil {
		usr = &user.User{Username: "unknown"}
	}

	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("xtest"),
			semconv.ServiceVersion(LocalVersion),
			semconv.EnduserID(usr.Username),
			attribute.Bool("sampling.exempt", true),
		),
	)
	return r
}

// StartTracer starts the OpenTelemetry tracer with the OTLP exporter.
func StartTracer() (func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	expo, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(otelEndpoint),
		otlptracehttp.WithRetry(otlptracehttp.RetryConfig{
			Enabled:         true,
			InitialInterval: initialInterval,
			MaxInterval:     maxInterval,
			MaxElapsedTime:  maxElapsedTime,
		}),
	)
	if err != nil {
		slog.Debug("Error creating OTLP exporter: ")
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(expo),
		sdktrace.WithResource(newResource()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := tp.ForceFlush(ctx); err != nil {
			slog.Debug("Failed to flush the trace provider: ")
		}
		if err := tp.Shutdown(ctx); err != nil {
			slog.Debug("Error shutting down tracer: ")
		}
		if err := expo.Shutdown(ctx); err != nil {
			slog.Debug("Failed to stop the exporter: ")
		}
	}, nil
}

```
