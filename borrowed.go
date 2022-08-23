package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createIcon() *widget.Icon {
	resc, _ := LoadResourceFromPath("https://github.com/arewabolu/book-points/blob/main/Assets/icons/bullet?raw=true")
	theme.NewThemedResource(resc)
	icon := widget.NewIcon(resc)

	return icon
}

func removeElementByIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func saveFunc(txt string, titleBind binding.String, noteBindings binding.ExternalStringList, w fyne.Window) {
	titleText, _ := titleBind.Get()
	writtenTitle, titleWriterErr := titleWriter(txt, titleText)

	if titleWriterErr != nil {
		anApp := app.New()
		errWind := anApp.NewWindow("error")
		dialog.ShowError(titleWriterErr, errWind)
	}

	noteList, err2 := noteBindings.Get()
	if err2 != nil {
		anApp := app.New()
		errWind := anApp.NewWindow("error")
		dialog.ShowError(err2, errWind)

	}
	write2Book(writtenTitle, noteList)
}
