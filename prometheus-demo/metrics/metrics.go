package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//There are three type of metrics in prometheus tool:
	//1. counter: used to count number of X event occured
	//2. guage: current value of X
	//3. histogram: how long or how big event ocuured
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestsInProgress = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_progress",
			Help: "Number of HTTP requests currently in progress",
		},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestsInProgress)
	prometheus.MustRegister(httpRequestDuration)
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

// For counter metric
func RecordRequest(method, endpoint, status string) {
	httpRequestsTotal.WithLabelValues(method, endpoint, status).Inc()
}

// For guage metric
func RequestStarted() {
	httpRequestsInProgress.Inc()
}

func RequestCompleted() {
	httpRequestsInProgress.Dec()
}

// For histogram metric
func StartTimer(method, endpoint string) *prometheus.Timer {
	return prometheus.NewTimer(httpRequestDuration.WithLabelValues(method, endpoint))

}
