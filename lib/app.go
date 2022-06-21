package prombackup

import (
	"fmt"
)

type App struct {
	Url      string
	SnapPath string
}

func (a *App) FullUrl() string {
	return (fmt.Sprintf("%s/%s", a.Url, a.SnapPath))
}
