package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
)

// ab -n 1000 -c 10 http://localhost:28080/live
func main() {
	runProfile()
	runGin()
}

func init() {
	runtime.SetBlockProfileRate(1)
}

func runProfile() {
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			panic(err)
		}
	}()
}

func runGin() {
	e := gin.Default()
	e.GET("/live", liveHandler)
	if err := e.Run(":28080"); err != nil {
		panic(err)
	}
}

func liveHandler(c *gin.Context) {
	mutexTaskService()
	c.JSON(http.StatusOK, map[string]int{"code": 0})
}

func mutexTaskService() {
	var m sync.Mutex
	datas := make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}
}
