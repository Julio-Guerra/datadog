package tracer

import (
	"github.com/Julio-Guerra/datadog/apm"
	"github.com/Julio-Guerra/datadog/appsec"
	"github.com/Julio-Guerra/datadog/debug"
	"github.com/Julio-Guerra/datadog/profiler"
)

func Hello() {
	appsec.Hello()
	profiler.Hello()
	apm.Hello()
	debug.Hello()
}
