package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/narharim/go-learning/prometheus-demo/handlers"
	"github.com/narharim/go-learning/prometheus-demo/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	metrics.InitMetrics()
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api/data", handlers.APIHandler)

	// Expose Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server starting on :8080")
	fmt.Println("Prometheus metrics available at /metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
