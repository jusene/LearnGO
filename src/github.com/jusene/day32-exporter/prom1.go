package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})

var hdFailures = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors",
	},
	[]string{"device"})

func init() {
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func main() {
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8111", nil))
}
