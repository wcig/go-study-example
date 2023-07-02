## 1.Counter

Counter: 只增不减的计数器.

服务运行后请求接口, 然后访问 `http://localhost:2112/metrics` 可以查看到如下指标: 

```
# HELP http_requests_total The HTTP requests total count.
# TYPE http_requests_total counter
http_requests_total{method="GET",path="/ping",service="go-app",status="200"} 5
http_requests_total{method="GET",path="/test",service="go-app",status="404"} 2
```

## 2.Guage

Gauge: 可增可减的仪表盘.

访问 `http://localhost:2112/metrics` 可以查看到如下指标: 

```
# HELP go_goroutines_number The go routine number.
# TYPE go_goroutines_number gauge
go_goroutines_number{service="go-app"} 8
```

## 3.Histogram

Histogram: 累积直方图.

访问 `http://localhost:2112/metrics` 可以查看到如下指标: 

```
# HELP http_request_duration_seconds The HTTP request durations in seconds.
# TYPE http_request_duration_seconds histogram
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="0.005"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="0.01"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="0.025"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="0.05"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="0.1"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="0.25"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="0.5"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="1"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="3"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="10"} 5
http_request_duration_seconds_bucket{method="GET",path="/ping",service="go-app",status="200",le="+Inf"} 5
http_request_duration_seconds_sum{method="GET",path="/ping",service="go-app",status="200"} 0.000231502
http_request_duration_seconds_count{method="GET",path="/ping",service="go-app",status="200"} 5
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="0.005"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="0.01"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="0.025"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="0.05"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="0.1"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="0.25"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="0.5"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="1"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="3"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="10"} 2
http_request_duration_seconds_bucket{method="GET",path="/test",service="go-app",status="404",le="+Inf"} 2
http_request_duration_seconds_sum{method="GET",path="/test",service="go-app",status="404"} 8.4e-07
http_request_duration_seconds_count{method="GET",path="/test",service="go-app",status="404"} 2
```

## 4.Summary

Summary: 摘要

访问 `http://localhost:2112/metrics` 可以查看到如下指标: 

```
# HELP http_request_duration_seconds_summary The HTTP request durations in seconds.
# TYPE http_request_duration_seconds_summary summary
http_request_duration_seconds_summary{method="GET",path="/ping",service="go-app",status="200",quantile="0.5"} 4.0678e-05
http_request_duration_seconds_summary{method="GET",path="/ping",service="go-app",status="200",quantile="0.9"} 7.2985e-05
http_request_duration_seconds_summary{method="GET",path="/ping",service="go-app",status="200",quantile="0.99"} 7.2985e-05
http_request_duration_seconds_summary_sum{method="GET",path="/ping",service="go-app",status="200"} 0.000231502
http_request_duration_seconds_summary_count{method="GET",path="/ping",service="go-app",status="200"} 5
http_request_duration_seconds_summary{method="GET",path="/test",service="go-app",status="404",quantile="0.5"} 3.7e-07
http_request_duration_seconds_summary{method="GET",path="/test",service="go-app",status="404",quantile="0.9"} 4.7e-07
http_request_duration_seconds_summary{method="GET",path="/test",service="go-app",status="404",quantile="0.99"} 4.7e-07
http_request_duration_seconds_summary_sum{method="GET",path="/test",service="go-app",status="404"} 8.4e-07
http_request_duration_seconds_summary_count{method="GET",path="/test",service="go-app",status="404"} 2
```
