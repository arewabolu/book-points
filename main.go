package main

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

type window struct {
	wind fyne.Window
}

func init() {
	w := new(window)
	err := makeDir()
	if err != nil {
		err = errors.New("unable to create new directory")
		dialog.NewError(err, w.wind)
	}
}

func main() {
	w := new(window)
	app := app.New()
	w.wind = app.NewWindow("BookTakes")

	w.wind.SetMaster()
	//wind.SetFullScreen(true)
	w.wind.Resize(fyne.NewSize(600, 600))

	fullWind := container.NewBorder(header(), nil, nil, nil, loadUI(w.wind))
	w.wind.SetContent(fullWind)
	w.wind.ShowAndRun()
}
