package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	goh "github.com/arewabolu/GoHaskell"
)

type bookinfo struct {
	name  string
	notes []string
}

func (b *bookinfo) remove(index int) {
	b.notes = goh.Pop(b.notes, index)
}

func (b *bookinfo) addNotes() {
	b.notes = append(b.notes, "")
}

func (b *bookinfo) changeLine(newText string) {
	if goh.Last(b.notes) != newText && goh.Last(b.notes) == "" {
		b.notes = goh.Put(b.notes, newText, len(b.notes)-1)
	}
}

func (b *bookinfo) saveNote(nwTitle string, w fyne.Window) {
	writtenTitle, titleWriterErr := titleWriter(b.name, nwTitle)

	if titleWriterErr != nil {
		anApp := app.New()
		errStr := fmt.Sprintf("error!\n unable to save file %s", nwTitle)
		errWind := anApp.NewWindow(errStr)
		dialog.ShowError(titleWriterErr, errWind)
	}

	write2Book(writtenTitle, b.notes)
}
