package main

import (
	"log"
	"net/http"

	"promundgrafana/internal/controller"
	"promundgrafana/internal/pkg/metric"
	"promundgrafana/internal/pkg/prometheus"
)

func main() {
	log.Println("server running on :8080")

	// Prometheus
	pro := prometheus.New(true)

	// Metric
	mtr := metric.New(pro.Registry())

	// Router
	rtr := http.NewServeMux()
	rtr.HandleFunc("/metrics", pro.Handler())
	rtr.HandleFunc("/api/v1/balances", controller.NewBalanceUpdate(mtr).Handle)

	// Server
	log.Fatalln(http.ListenAndServe(":8080", rtr))
}