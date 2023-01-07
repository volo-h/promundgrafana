package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

func balanceGauge() prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "client",
		Name:      "balance_gauge",
		Help:      "Current balance",
	})
}
