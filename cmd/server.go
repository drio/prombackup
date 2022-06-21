package main

import (
	prombackup "github.com/drio/prombackup/lib"
	"log"
)

func main() {
	app := &prombackup.App{
		SnapUrl:    "http://localhost:9090",
		SnapPath:   "api/v1/admin/tsdb/snapshot",
		ListenPort: ":8080",
		S3Region:   "us-east-1",
		S3Bucket:   "drio-prom-timeseries-snap",
		S3ACL:      "public-read",
	}

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

	//app.Run()

	tarBall, err := prombackup.MakeTarBall("data/snapshots/20220621T120952Z-5272c333caf89e5d")
	if err != nil {
		log.Println("Ups, problems making tarball: ", err)
	} else {
		log.Println("tarball: ", tarBall)
	}

	err = app.UploadFile(tarBall)
	if err != nil {
		log.Println("Error uploading tarball to S3: ", err)
	} else {
		log.Println("Success uploading tarball to S3")
	}
}
