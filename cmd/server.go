package main

import (
	prombackup "github.com/drio/prombackup/lib"
)

func main() {
	app := &prombackup.App{
		Url:      "http://localhost:9090",
		SnapPath: "api/v1/admin/tsdb/snapshot",
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

	app.Run()
}
