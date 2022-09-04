package main

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

func main() {
	app := app.New()
	wind := app.NewWindow("BookTakes")

	wind.SetMaster()
	//wind.SetFullScreen(true)
	wind.Resize(fyne.NewSize(600, 600))

	err := makeDir()
	if err != nil {
		err = errors.New("unable to create new directory")
		dialog.NewError(err, wind)
	}

	fullWind := container.NewBorder(header(), nil, nil, nil, loadUI(wind))
	wind.SetContent(fullWind)
	wind.ShowAndRun()
}
