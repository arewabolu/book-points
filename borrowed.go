package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

/*func createIcon() *widget.Icon {
	resc, err := LoadResourceFromPath("./Assets/icons/bullet")
	if err != nil {
		panic(err)
	}
	theme.NewThemedResource(resc)
	icon := widget.NewIcon(resc)

	return icon
}
*/
//func createLockIcon() *widget.Icon {
//	resc, err := LoadResourceFromPath("./Assets/icons/lock.png")
//	if err != nil {
//		panic(err)
//	}
//	theme.NewThemedResource(resc)
//	icon := widget.NewIcon(resc)
//
//	return icon
//}

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
