package tracer

import (
	"github.com/cloudwego/kitex/server"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"io"
	"time"

	"github.com/uber/jaeger-client-go/config"

	trace "github.com/kitex-contrib/tracer-opentracing"
)

func newJaegerTracer(serviceName, agentHostPort string) (server.Suite, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	opentracing.InitGlobalTracer(tracer)
	return trace.NewDefaultServerSuite(), closer, err
}

func Init(serviceName string) (server.Suite, io.Closer) {
	tracerSuite, closer, err := newJaegerTracer(serviceName, "127.0.0.1:6831")
	if err != nil {
		panic(err)
	}
	return tracerSuite, closer
}
