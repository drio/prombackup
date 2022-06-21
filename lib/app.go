package prombackup

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	Url      string
	SnapPath string
}

func (app *App) FullUrl() string {
	return (fmt.Sprintf("%s/%s", app.Url, app.SnapPath))
}

func (app *App) Run() {
	http.HandleFunc("/backup", app.HandleSnapReq)
	http.Handle("/metrics", promhttp.Handler())
	port := ":8080"
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
