package prombackup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SnapResponse struct {
	Status string   `json:"status"`
	Data   SnapData `json:"data"`
}

type SnapData struct {
	Name string `json:"name"`
}

var runningSnapshot = false

func makeRequest(url string) ([]byte, error) {
	res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte("")))
	if err != nil {
		log.Println("Error making post request:", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading returned post body:", err)
		return nil, err
	}

	return []byte(body), err
}

func (app *App) MakeTarBall(sourceDir string) (string, error) {
	outputFile := app.TarBallName
	log.Printf("Trying to make tarball %s on snapshot %s", outputFile, sourceDir)

	// tar + gzip
	var buf bytes.Buffer
	err := compress(sourceDir, &buf)
	if err != nil {
		return "", err
	}

	// write the .tar.gz
	fileToWrite, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, os.FileMode(0600))
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(fileToWrite, &buf); err != nil {
		return "", err
	}

	return outputFile, nil
}

func (app *App) CreateSnapShot() (string, error) {
	log.Println("Making snapshot request:", app.FullUrl())
	body, err := makeRequest(app.FullUrl())
	if err != nil {
		return "", err
	}

	var sr = new(SnapResponse)
	err = json.Unmarshal(body, &sr)
	if err != nil {
		log.Println("Error unmarshalling Snap Response", err)
		return "", err
	}

	return sr.Data.Name, nil
}

func (app *App) RunSnapShot() {
	func() {
		runningSnapshot = true
		defer app.cleanUp()
		snapName, err := app.CreateSnapShot()
		if snapName == "" {
			log.Println("Snapname is empty", err)
			return
		}
		if err != nil {
			log.Println("Error creating snapshot:", err)
			return
		}
		log.Println("Snap dir created:", snapName)

		tarBall, err := app.MakeTarBall(fmt.Sprintf("%s/%s", app.SnapDir, snapName))
		if err != nil {
			log.Println("Error making tarball:", err)
			return
		}

		snapSize, err := FileSize(tarBall)
		if err != nil {
			log.Println("Error extracting tarball size:", err)
			return
		}
		log.Println("Tarball snapshot size:", snapSize)

		err = app.UploadFile(tarBall)
		if err != nil {
			log.Println("Could not upload tarbal to S3: ", err)
			return
		}
		log.Println("Success uploading tarball to S3")
		BackupSize.Add(float64(snapSize))

		log.Printf("Snapshot pipeline completed: %s %d", snapName, snapSize)
	}()
}

func (app *App) cleanUp() {
	log.Println("Starting cleanup")
	runningSnapshot = false

	if _, err := os.Stat(app.SnapDir); err == nil {
		log.Println("Cleaning up snapdir:", app.SnapDir)
		err := os.RemoveAll(app.SnapDir)
		if err != nil {
			log.Println("Error cleaning up SnapDir:", err)
		}
	}

	if _, err := os.Stat(app.TarBallName); err == nil {
		log.Println("Cleaning up tarball:", app.TarBallName)
		// File exists
		err = os.RemoveAll(app.TarBallName)
		if err != nil {
			log.Println("Error cleaning up Tarball:", err)
		}
	}
}
