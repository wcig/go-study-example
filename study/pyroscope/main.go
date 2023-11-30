package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grafana/pyroscope-go"
)

// ab -n 100 -c 10 http://localhost:28080/live
func main() {
	runPyroscope()
	runGin()
}

func runPyroscope() {
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
	})
	if err != nil {
		panic(err)
	}
	_ = profiler
}

func runGin() {
	e := gin.Default()
	e.GET("/live", liveHandler)
	if err := e.Run(":28080"); err != nil {
		panic(err)
	}
}

func liveHandler(c *gin.Context) {
	cpuTaskService()
	c.JSON(http.StatusOK, map[string]int{"code": 0})
}

func cpuTaskService() {
	for i := 0; i < 1e9; i++ {
		//
	}
}
