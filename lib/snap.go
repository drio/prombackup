package prombackup

import (
	"bytes"
	"encoding/json"
	"fmt"
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

/*
  At x time during the day:
   - post to prometheus and get ID
   - send directory to S3 or BB
     * if ok:
       - measure the directory size and update backupSize
       - update metric to backpSize
       - remove the snapshot directory
     * else:
       - update metric to -1
   - After x minutes, update the metric value back to zero
*/
func (app *App) HandleSnapReq(w http.ResponseWriter, req *http.Request) {
	pName, err := app.CreateSnapShot()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Could not create snapshot")
	} else {
		fmt.Fprintf(w, "ok: %s", *pName)
	}
}
