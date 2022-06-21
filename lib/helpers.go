package prombackup

import (
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
