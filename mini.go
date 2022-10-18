package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func CreateNewLineButton(s string, b binding.ExternalStringList, l *widget.List) *widget.Button {
	return widget.NewButtonWithIcon(s, theme.ContentAddIcon(), func() {
		b.Append("")
	})
}

func CreateNewList(notes []string, bind binding.ExternalStringList) *widget.List {
	listing := widget.NewListWithData(
		bind,
		func() fyne.CanvasObject {
			noteEntry := widget.NewEntry()

			//var ev *fyne.KeyEvent

			delButn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				position := slices.Index(notes, noteEntry.Text)
				nwList := removeElementByIndex(notes, position)
				bind.Set(nwList)
			})

			delButn.Resize(fyne.NewSize(10, 10))
			moveButn := widget.NewButtonWithIcon("", theme.ListIcon(), func() {
			})
			cont := container.NewBorder(nil, nil, moveButn, delButn, noteEntry)
			return cont
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*fyne.Container).Objects[0].(*widget.Entry).Bind(di.(binding.String))
		},
	)
	return listing
}
