package main

import (
	prombackup "github.com/drio/prombackup/lib"
	"log"
)

func main() {
	/*
		app := &prombackup.App{
			SnapUrl:    "http://localhost:9090",
			SnapPath:   "api/v1/admin/tsdb/snapshot",
			ListenPort: ":8080",
		}
	*/

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
	out, err := prombackup.MakeTarBall("data/snapshots/20220621T120952Z-5272c333caf89e5d")
	if err != nil {
		log.Println("Ups ", err)
	} else {
		log.Println("All good: ", out)
	}
}
