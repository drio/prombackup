package main

import (
	"flag"
	prombackup "github.com/drio/prombackup/lib"
)

func main() {
	serverUrl := flag.String("server", "http://localhost:9090", "prometheus server url (including port)")
	snapDir := flag.String("snapdir", "./prometheus/data/snapshots", "snapshots directory path")
	secondsToZero := flag.Int("zeroSeconds", 60, "wait secondsToZero to set metric back to zero after snapshot")
	tarballName := flag.String("tarballName", "prom-snapshot.tar.gz", "tarball file name")
	listenPort := flag.Int("listenPort", 8080, "port to listen to")
	s3Region := flag.String("s3region", "us-east-1", "AWS s3 region")
	s3Bucket := flag.String("s3bucket", "drio-prom-timeseries-snap", "S3 bucket name")

	flag.Parse()

	app := &prombackup.App{
		ServerUrl:     *serverUrl,
		SnapPath:      "api/v1/admin/tsdb/snapshot",
		SnapDir:       *snapDir,
		SecondsToZero: *secondsToZero,
		TarBallName:   *tarballName,
		ListenPort:    *listenPort,
		S3Region:      *s3Region,
		S3Bucket:      *s3Bucket,
	}

	app.Run()
}
