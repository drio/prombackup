package prombackup

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	backupSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "prom_backup_size",
		Help: "The size of the latest snapshot",
	})
)

func init() {
	prometheus.MustRegister(backupSize)
	backupSize.Set(0.2)
}
