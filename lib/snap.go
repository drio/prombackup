package prombackup

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func (app *App) CreateSnapShot() (*string, error) {
	body, err := makeRequest(app.FullUrl())
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
