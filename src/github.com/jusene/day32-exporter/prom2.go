package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

type fooCollector struct {
	fooMetric *prometheus.Desc
	barMetric *prometheus.Desc
}

func newFooCollector() *fooCollector {
	m1 := make(map[string]string)
	m1["env"] = "prod"
	v := []string{"hostname", "local"}
	return &fooCollector{
		fooMetric: prometheus.NewDesc("fff_metrics", "Show metrics a for mysql", nil, nil),
		barMetric: prometheus.NewDesc("bbb_metrics", "Show metrics a bar occu", v, m1),
	}
}

func (collect *fooCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collect.barMetric
	ch <- collect.fooMetric

}

func (collect *fooCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64

	if 1 == 1 {
		metricValue = 1
	}

	ch <- prometheus.MustNewConstMetric(collect.fooMetric, prometheus.GaugeValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collect.barMetric, prometheus.CounterValue, metricValue, "kk", "sasa")
}

func main() {
	foo := newFooCollector()
	prometheus.MustRegister(foo)
	log.Info("beging to server on Port: 18080")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":18080", nil))
}
