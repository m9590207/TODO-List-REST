package global

import "github.com/opentracing/opentracing-go"

//服務的全域變數
var (
	Tracer opentracing.Tracer
)
