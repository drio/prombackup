package prombackup

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	BackupSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "prom_backup_size",
		Help: "The size in bytes of the latest snapshot",
	})
)

func init() {
	prometheus.MustRegister(BackupSize)
	BackupSize.Set(0)
}
