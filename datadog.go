package tracer

import (
	"github.com/Julio-Guerra/datadog/tracer"
	"github.com/Julio-Guerra/datadog/appsec"
	"github.com/Julio-Guerra/datadog/debug"
	"github.com/Julio-Guerra/datadog/profiler"
)

func Hello() {
	appsec.Hello()
	profiler.Hello()
	tracer.Hello()
	debug.Hello()
}
