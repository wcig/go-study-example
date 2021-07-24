package gin

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func TestPrometheus(t *testing.T) {
	go InitMetric()

	router := gin.Default()
	router.Use(MetricHandleFunc)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.Run(":28080")
}

var (
	HTTPReqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        "http_access_metric",
			Help:        "The HTTP request latencies in seconds.",
			Buckets:     nil,
			ConstLabels: prometheus.Labels{"service": "go-app"},
		},
		[]string{"method", "api", "status"},
	)
)

// prometheus监控
func InitMetric() {
	// 记录请求信息
	prometheus.MustRegister(HTTPReqDuration)

	// 记录基础信息
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func MetricHandleFunc(c *gin.Context) {
	start := time.Now()
	c.Next()
	duration := float64(time.Since(start).Milliseconds())
	HTTPReqDuration.With(prometheus.Labels{
		"method": c.Request.Method,
		"api":    c.Request.URL.Path,
		"status": strconv.Itoa(c.Writer.Status()),
	}).Observe(duration)
}
