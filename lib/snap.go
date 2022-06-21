package prombackup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type SnapResponse struct {
	Status string   `json:"status"`
	Data   SnapData `json:"data"`
}

type SnapData struct {
	Name string `json:"name"`
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

func (app *App) MakeTarBall(sourceDir string) (string, error) {
	outputFile := app.TarBallName
	log.Printf("Trying to make tarball %s on snap %s", outputFile, sourceDir)
	_, err := exec.Command("tar", "-zcf", outputFile, sourceDir).CombinedOutput()
	if err != nil {
		return "", err
	}
	return outputFile, nil
}

func (app *App) CreateSnapShot() (string, error) {
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
		snapName, err := app.CreateSnapShot()
		if err != nil {
			log.Println("Error creating snapshot:", err)
			return
		}
		log.Println("Snap dir created:", snapName)

		tarBall, err := app.MakeTarBall(fmt.Sprintf("%s/%s", app.SnapDir, snapName))
		if err != nil {
			log.Println("Error making tarball:", err)
			BackupSize.Set(-1)
			return
		}
		defer app.cleanUp()

		snapSize, err := FileSize(tarBall)
		if err != nil {
			log.Println("Error extracting tarball size:", err)
			BackupSize.Set(-1)
			return
		}
		log.Println("Tarball snapshot size:", snapSize)

		err = app.UploadFile(tarBall)
		if err != nil {
			log.Println("Could not upload tarbal to S3: ", err)
			BackupSize.Set(-1)
			return
		}
		log.Println("Success uploading tarball to S3")
		BackupSize.Set(float64(snapSize))

		go func() {
			time.Sleep(time.Duration(app.SecondsToZero) * time.Second)
			log.Println("Setting backupsize metric back to zero")
			BackupSize.Set(0)
		}()

		log.Printf("Snapshot pipeline completed: %s %d", snapName, snapSize)
	}()
}

func (app *App) cleanUp() {
	err := os.RemoveAll(app.SnapDir)
	if err != nil {
		log.Println("Error cleaning up SnapDir: ", err)
	}

	if _, err := os.Stat(app.TarBallName); err == nil {
		// File exists
		err = os.RemoveAll(app.TarBallName)
		if err != nil {
			log.Println("Error cleaning up Tarball: ", err)
		}
	}
}
