package prombackup

import (
	"fmt"
)

type App struct {
	Url      string
	SnapPath string
}

func (app *App) FullUrl() string {
	return (fmt.Sprintf("%s/%s", app.Url, app.SnapPath))
}
