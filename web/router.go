package web

import (
	"I2Oprobe/internal/g"
	"I2Oprobe/internal/log"
	"I2Oprobe/prom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	metricPath = "/metrics"
	nameSpace = "I2ONET"
)

func Start() {

	log.Info("Http start")
	metrics := prom.NewMetrics(nameSpace)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET(metricPath, gin.WrapH(promhttp.HandlerFor(registry, promhttp.HandlerOpts{})))

	router.Run(*g.ListenAndPort)
}

