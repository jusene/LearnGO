## GO语言 Prometheus

### 指标类型

- Counter(累加指标)
- Gauge(测量指标)
- Summary(概略图)
- Histogram(直方图)

Counter是一种累加指标数据，这个值随着时间只会逐渐的增加。比如完成的总任务数量，运行错误发生的总次数。

Gauge代表单一数据，数据可增减，如CPU使用情况

Histogram和Summary，使用基于采样的模式，总的响应时间低于300ms的占比，或者查询95%用户查询的门限值对应的响应时间是多少。 使用Histogram和Summary指标的时候同时会产生多组数据，_count代表了采样的总数，_sum则代表采样值的和。 _bucket则代表了落入此范围的数据。如果需要聚合数据，可以使用histogram. 并且如果对于分布范围有明确的值的情况下（比如300ms），也可以使用histogram。但是如果仅仅是一个百分比的值（比如上面的95%），则使用Summary

- 源码接口
```
type Collector interface {
	// Describe sends the super-set of all possible descriptors of metrics
	// collected by this Collector to the provided channel and returns once
	// the last descriptor has been sent. The sent descriptors fulfill the
	// consistency and uniqueness requirements described in the Desc
	// documentation.
	//
	// It is valid if one and the same Collector sends duplicate
	// descriptors. Those duplicates are simply ignored. However, two
	// different Collectors must not send duplicate descriptors.
	//
	// Sending no descriptor at all marks the Collector as “unchecked”,
	// i.e. no checks will be performed at registration time, and the
	// Collector may yield any Metric it sees fit in its Collect method.
	//
	// This method idempotently sends the same descriptors throughout the
	// lifetime of the Collector. It may be called concurrently and
	// therefore must be implemented in a concurrency safe way.
	//
	// If a Collector encounters an error while executing this method, it
	// must send an invalid descriptor (created with NewInvalidDesc) to
	// signal the error to the registry.
	Describe(chan<- *Desc)
	// Collect is called by the Prometheus registry when collecting
	// metrics. The implementation sends each collected metric via the
	// provided channel and returns once the last metric has been sent. The
	// descriptor of each sent metric is one of those returned by Describe
	// (unless the Collector is unchecked, see above). Returned metrics that
	// share the same descriptor must differ in their variable label
	// values.
	//
	// This method may be called concurrently and must therefore be
	// implemented in a concurrency safe way. Blocking occurs at the expense
	// of total performance of rendering all registered metrics. Ideally,
	// Collector implementations support concurrent readers.
	Collect(chan<- Metric)
}
```

```go
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
```