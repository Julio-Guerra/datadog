package tracer

import (
	"github.com/Julio-Guerra/tracer/apm"
	"github.com/Julio-Guerra/tracer/appsec"
	"github.com/Julio-Guerra/tracer/debug"
	"github.com/Julio-Guerra/tracer/profiler"
)

func Hello() {
	appsec.Hello()
	profiler.Hello()
	apm.Hello()
	debug.Hello()
}
