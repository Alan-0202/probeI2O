package prom

import (
	"I2Oprobe/internal/g"
	"I2Oprobe/internal/model"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"sync"
)

var (
	ms = make(map[string] *prometheus.Desc)
)

type Metrics struct {
	metrics map[string] *prometheus.Desc
	mutex sync.Mutex
}

func NewGlobalMetric(namespace string, metricName string, help string, labels []string) *prometheus.Desc{
	return prometheus.NewDesc(namespace + "_" + metricName, help, labels, nil)
}

func NewMetrics(ns string) *Metrics {
	ms["check"] = NewGlobalMetric(ns, "bad_request", "bad http request", []string{"address"})

	//Ip     string
	//Port   float64
	//Meg    string
	//Value  float64
	//Shop   string
	//Ration float64
	ms["result"] = NewGlobalMetric(ns, "result_request", "get ration", []string{"ip", "port", "Meg", "shop", "ration"})
	return &Metrics{
		metrics: ms,
	}
}

func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}


func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	result := func(k,v interface{}) bool {

		valFloat64, err := strconv.ParseFloat(v.(string),64 )
		if err != nil {
			return false
		}

		ch <- prometheus.MustNewConstMetric(c.metrics["check"], prometheus.GaugeValue, valFloat64,
			k.(string))

		return true
	}
	g.MetricsReqMap.Range(result)

	goodRes := func(k interface{},val interface{}) bool {
		for _,v := range val.([]model.MetricsOpts) {
			ch <- prometheus.MustNewConstMetric(c.metrics["result"], prometheus.GaugeValue,
				float64(v.Value),
				v.Ip,
				strconv.Itoa(v.Port),
				v.Meg,
				v.Shop,
				strconv.Itoa(v.Ration))
		}
		return true
	}

	g.MetricsMap.Range(goodRes)

}
