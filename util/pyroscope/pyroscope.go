package pyroscope

import (
	"os"

	"github.com/grafana/pyroscope-go"
)

func StartPyroscope() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	profiler, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "simple.golang.app",
		ServerAddress:   "http://localhost:4040",
		Logger:          pyroscope.StandardLogger,

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
		Tags: map[string]string{
			"hostname": hostname,
		},
	})
	if err != nil {
		panic(err)
	}
	_ = profiler
}
