package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	//jaeger client的設定項目
	cfg := &config.Configuration{
		ServiceName: serviceName,
		//設定採樣的類型,固定取樣對所有資料都取樣
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		//是否啓用logging reporter 更新緩衝區的頻率, 上報的agent位址
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	//設定全域的tracer物件,如果不設定會致上下文無法產生正確的span
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
