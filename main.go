package main

/*
  1. periodically send a post to the prometheus server
  2. send the data to backblaze (https://github.com/kurin/blazer) or S3.
  3. write a prom exporter
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
)

type SnapResponse struct {
	Status string   `json:"status"`
	Data   SnapData `json:"data"`
}

type SnapData struct {
	Name string `json:"name"`
}

var (
	backupSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "progo_backup_done",
		Help: "The size of the latest snapshot",
	})
)

func init() {
	prometheus.MustRegister(backupSize)
}

func makeRequest(url string) ([]byte, error) {
	res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte("")))
	if err != nil {
		log.Println("Error making post request: ", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading returned post body: ", err)
		return nil, err
	}

	return []byte(body), err
}

func createSnapShot() (*string, error) {
	body, err := makeRequest("http://localhost:9090/api/v1/admin/tsdb/snapshot")
	if err != nil {
		return nil, err
	}

	var sr = new(SnapResponse)
	err = json.Unmarshal(body, &sr)
	if err != nil {
		log.Println("Error unmarshalling Snap Response", err)
		return nil, err
	}

	return &sr.Data.Name, nil
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
	http.HandleFunc("/backup", handleBackup)
	backupSize.Set(0)

	go func() {
		max := 30
		min := 0
		for {
			r := rand.Intn(max-min) + min
			backupSize.Set(float64(r))
			time.Sleep(1)
		}
	}()

	pName, err := createSnapShot()
	if err != nil {
		log.Println(err)
	}
	log.Println(*pName)

	//http.Handle("/metrics", promhttp.Handler())
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
