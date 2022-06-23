package prombackup

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	BackupSize = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "prom_backup_size_count",
			Help: "The size in bytes of the latest snapshot",
		},
	)
)

func init() {
	prometheus.MustRegister(BackupSize)
}
