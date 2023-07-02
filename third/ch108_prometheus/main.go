package main

import (
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go initMetric()
	runServer()
}

func runServer() {
	router := gin.Default()
	router.Use(MetricHandleFunc)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	if err := router.Run(":28080"); err != nil {
		panic(err)
	}
}

const (
	serviceName = "go-app"
)

var (
	serviceNameLabel = prometheus.Labels{"service": serviceName}

	// Counter: 只增不减的计数器
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "http_requests_total",
			Help:        "The HTTP requests total count.",
			ConstLabels: serviceNameLabel,
		},
		[]string{"method", "path", "status"},
	)

	// Gauge: 可增可减的仪表盘
	goroutinesNumber = prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "go_goroutines_number",
			Help:        "The go routine number.",
			ConstLabels: serviceNameLabel,
		},
		func() float64 {
			return float64(runtime.NumGoroutine())
		},
	)

	// Histogram: 累积直方图
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        "http_request_duration_seconds",
			Help:        "The HTTP request durations in seconds.",
			Buckets:     []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 3, 10},
			ConstLabels: serviceNameLabel,
		},
		[]string{"method", "path", "status"},
	)

	// Summary: 摘要
	httpRequestDurationSummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:        "http_request_duration_seconds_summary",
			Help:        "The HTTP request durations in seconds.",
			Objectives:  map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			ConstLabels: serviceNameLabel,
		},
		[]string{"method", "path", "status"},
	)
)

func initMetric() {
	collectors := []prometheus.Collector{
		httpRequestsTotal,
		goroutinesNumber,
		httpRequestDuration,
		httpRequestDurationSummary,
	}
	prometheus.MustRegister(collectors...)

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":2112", nil); err != nil {
		panic(err)
	}
}

func MetricHandleFunc(c *gin.Context) {
	start := time.Now()
	c.Next()
	duration := float64(time.Since(start).Nanoseconds()) / 1e9

	httpRequestsTotal.With(prometheus.Labels{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"status": strconv.Itoa(c.Writer.Status()),
	}).Inc()

	httpRequestDuration.With(prometheus.Labels{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"status": strconv.Itoa(c.Writer.Status()),
	}).Observe(duration)

	httpRequestDurationSummary.With(prometheus.Labels{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"status": strconv.Itoa(c.Writer.Status()),
	}).Observe(duration)
}
