package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	prombackup "github.com/drio/prombackup/lib"
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

/*
  At x time during the day:
   - post to prometheus and get ID
   - send directory to S3 or BB
     * if ok:
       - measure the directory size and update backupSize
       - update metric to backpSize
       - remove the snapshot directory
     * else:
       - update metric to -1
   - After x minutes, update the metric value back to zero
*/
func handleBackup(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {

	app := &prombackup.App{
		Url:      "http://localhost:9090",
		SnapPath: "api/v1/admin/tsdb/snapshot",
	}

	http.HandleFunc("/backup", handleBackup)
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

	pName, err := app.CreateSnapShot()
	if err != nil {
		log.Println("Error creating snapshot: ", err)
	} else {
		log.Println(*pName)
	}

	//http.Handle("/metrics", promhttp.Handler())
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
