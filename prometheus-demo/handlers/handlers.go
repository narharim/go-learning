package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/narharim/go-learning/prometheus-demo/metrics"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	timer := metrics.StartTimer(r.Method, "/")
	metrics.RequestStarted()
	defer metrics.RequestCompleted()
	defer timer.ObserveDuration()

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	w.Write([]byte("Hello, Prometheus!"))

	metrics.RecordRequest(r.Method, "/", "200")
}

func APIHandler(w http.ResponseWriter, r *http.Request) {

	timer := metrics.StartTimer(r.Method, "/api/data")
	metrics.RequestStarted()
	defer metrics.RequestCompleted()
	defer timer.ObserveDuration()

	//Do some random work
	processingTime := time.Duration(rand.Intn(200)) * time.Millisecond
	time.Sleep(processingTime)

	//Randomly return an error
	if rand.Float64() < 0.1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		metrics.RecordRequest(r.Method, "/api/data", "500")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "success", "data": "some data"}`))
	metrics.RecordRequest(r.Method, "/api/data", "200")
}
