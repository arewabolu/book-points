package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func sublineEntry() *widget.Entry {
	subButn := widget.NewEntry()
	subButn.Move(fyne.NewPos(100, 30))
	return subButn
}
func newLnEntry() *widget.Entry { return widget.NewEntry() }

func oneAdd() *widget.Button {
	b := widget.NewButton("NewLine", func() {})
	return b
}

func createNewList(notes []string, listEnt *widget.Entry, bind binding.ExternalStringList) *widget.List {
	//	var intt *int
	/*alist := widget.NewList(
		func() int {
			return len(notes)
		},
		func() fyne.CanvasObject {
			noteEntry := widget.NewEntry()
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
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			ent := co.(*fyne.Container).Objects[0].(*widget.Entry)
			if strings.Contains(notes[lii], "\t") {
				ent = sublineEntry()
				ent.SetText(notes[lii])
			}
			ent.SetText(notes[lii])

		},
	)*/

	listing := widget.NewListWithData(
		bind,
		func() fyne.CanvasObject {

			delButn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				position := slices.Index(notes, listEnt.Text)
				nwList := removeElementByIndex(notes, position)
				bind.Set(nwList)
			})

			delButn.Resize(fyne.NewSize(10, 10))
			moveButn := widget.NewButtonWithIcon("", theme.ListIcon(), func() {
			})
			cont := container.NewBorder(nil, nil, moveButn, delButn, listEnt)
			return cont
		},
		func(di binding.DataItem, co fyne.CanvasObject) {

			line, _ := di.(binding.String).Get()
			if !strings.Contains(line, "\t") {
				entry := co.(*fyne.Container).Objects[0].(*widget.Entry)
				entry = sublineEntry()
				entry.Bind(binding.BindString(&line))
			}
			//	entry = newLnEntry()
			//			entry.SetText(line)

		},
	)
	return listing
}
