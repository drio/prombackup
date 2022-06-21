package main

import (
	prombackup "github.com/drio/prombackup/lib"
)

func main() {
	app := &prombackup.App{
		SnapUrl:     "http://localhost:9090",
		SnapPath:    "api/v1/admin/tsdb/snapshot",
		SnapDir:     "./data/snapshots",
		TarBallName: "prom-snapshot.tar.gz",
		ListenPort:  ":8080",
		S3Region:    "us-east-1",
		S3Bucket:    "drio-prom-timeseries-snap",
		S3ACL:       "public-read",
	}

	app.Run()
}
