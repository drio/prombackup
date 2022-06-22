package prombackup

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	ServerUrl     string
	SnapPath      string
	SnapDir       string
	SecondsToZero int // After SecondsToZero set the metric back to zero bytes
	TarBallName   string
	ListenPort    int
	S3Region      string
	S3Bucket      string
}

func (app *App) FullUrl() string {
	return (fmt.Sprintf("%s/%s", app.ServerUrl, app.SnapPath))
}

func (app *App) Run() {
	http.HandleFunc("/snap", func(w http.ResponseWriter, r *http.Request) {
		go app.RunSnapShot()
		fmt.Fprintf(w, "ok")
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listening on port", app.ListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.ListenPort), nil))
}
