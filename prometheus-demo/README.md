### prometheus-demo  
- a simple project that helps you understand how Prometheus works. I had fun building this
- focuses on three types of metrics: counter, gauge, and histogram

### try it out  
- run `go run main.go`
- trigger `http://localhost:8080` or `http://localhost:8080/api/data`
- Prometheus will collect the metrics, which you can view at `http://localhost:8080/metrics`
- To see it in the UI, run:  
  ```sh
    docker run --network=host -v ./prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
      ```  
- you can query using PromQL in browser `http:localhost:9090` and have fun! :)
