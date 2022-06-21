package prombackup

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func FileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	size := fi.Size()
	return size, nil
}

func commonError(msg string, w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	fmt.Fprintf(w, "Could not create snapshot")
	log.Println(msg, err)
}
