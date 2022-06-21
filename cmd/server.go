package main

import (
	"log"
	"net/http"

	prombackup "github.com/drio/prombackup/lib"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	backupSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "progo_backup_done",
		Help: "The size of the latest snapshot",
	})
)

func init() {
	prometheus.MustRegister(backupSize)
}

func main() {
	app := &prombackup.App{
		Url:      "http://localhost:9090",
		SnapPath: "api/v1/admin/tsdb/snapshot",
	}

	http.HandleFunc("/backup", app.HandleSnapReq)
	backupSize.Set(0)

	/*
		go func() {
			max := 30
			min := 0
			for {
				r := rand.Intn(max-min) + min
				backupSize.Set(float64(r))
				time.Sleep(1)
			}
		}()
	*/

	http.Handle("/metrics", promhttp.Handler())
	port := ":8080"
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
