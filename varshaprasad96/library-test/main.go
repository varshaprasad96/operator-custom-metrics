package main

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	metricslib "github.com/varshaprasad96/operator-custom-metrics"

	log "github.com/sirupsen/logrus"
)

// Metrics endpoint and path which is to be used to expose metrics.
const (
	metricsEndPoint = "8080"
	metricsPath     = "/metrics"
)

// Metric variables are are to be collected.
var (
	opsProcessed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events Test",
	})
	gaugeTest = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_total_new",
		Help: "The total number of processed events Test Again!",
	})
)

// RecordMetrics updates the values of the metrics which are to be collected.
func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			gaugeTest.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {

	prTest := metricslib.NewBuilder().
		WithPort(metricsEndPoint).
		WithCollectors(opsProcessed).
		GetConfig()

	if err := metricslib.ConfigureMetrics(context.TODO(), *prTest); err != nil {
		log.Error(err, "Fail")
	}
}
