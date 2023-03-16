package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func NewRightSide(b *bookinfo, w fyne.Window) fyne.CanvasObject {
	title, binding := titleEntry(b.name)

	oneAdd := oneAdd()
	butnLine := container.NewBorder(nil, nil, oneAdd, nil)
	lists := container.NewVBox()
	saveButn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		titleText, err := binding.Get()
		if err != nil {
			anApp := app.New()
			errWind := anApp.NewWindow("error")
			dialog.ShowError(err, errWind)
		}
		b.saveNote(titleText, w)
	})

	fLine := container.NewVBox(title, saveButn)
	topQuater := container.NewVBox(fLine, butnLine)
	rightHand2 := container.NewBorder(topQuater, nil, nil, nil)
	rightScroll := container.NewScroll(rightHand2)

	for _, item := range b.notes {
		lists.Add(minibox(b, item, rightHand2))
		rightHand2.Add(lists)
	}
	oneAdd.OnTapped = func() {
		b.addNotes()
		lists.Add(minibox(b, "", rightHand2))
		rightHand2.Add(lists)
		rightHand2.Refresh()
		rightScroll.Refresh()
	}
	DictionButn := widget.NewButton("Dictionaries", func() {})
	//sublineButn := widget.NewButton("Subline", func() {})

	//saveButn.Resize(fyne.NewSize(30, 30))

	//

	//widget.NewEntryWithData()

	DictionButn2 := widget.NewButtonWithIcon("Back to note", theme.LogoutIcon(), func() {
		rightHand2.RemoveAll()
		butnLine.RemoveAll()
		butnLine.Add(oneAdd)
		butnLine.Add(DictionButn)
		rightHand2.Add(topQuater)
		//	rightHand2.Add(listing)
	})

	DictionButn.OnTapped = func() {
		//	saveFunc(names[ID], titleBind, noteBindings, w)
		rightHand2.RemoveAll()
		butnLine.RemoveAll()
		butnLine.Add(DictionButn2)
		x := widget.NewMultiLineEntry()
		rightHand2.Add(topQuater)
		rightHand2.Add(x)
		rightScroll.Refresh()
	}

	return rightScroll
}

func minibox(b *bookinfo, text string, contnr *fyne.Container) *fyne.Container {
	listEnt, _ := titleEntry(text)
	delButn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		position := slices.Index(b.notes, listEnt.Text)
		fmt.Println(listEnt.Text)
		fmt.Println(position)
		if position == -1 {
			return
		}
		b.remove(position)
		contnr.Refresh()
	})

	delButn.Resize(fyne.NewSize(10, 10))

	cont := container.NewBorder(nil, nil, nil, delButn, listEnt)
	return cont
}
