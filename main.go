package main

import (
	"errors"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//text-header for application
func header() fyne.CanvasObject {
	rect := &canvas.Rectangle{
		FillColor:   color.White,
		StrokeColor: color.Black,
		StrokeWidth: 1,
	}

	width := rect.MinSize().Width

	rect.Move(fyne.NewPos(5, 5))
	rect.Resize(fyne.NewSize(width, 90))

	topic := canvas.NewText("Books", color.Black)
	topic.Alignment = fyne.TextAlignCenter
	//using maxLayout to stack text on top of rectangle
	header := container.New(layout.NewMaxLayout(), rect, topic)

	return header
}

//Loads when note buuton is clicked
func rightSide(txt, dir string) fyne.CanvasObject {
	title, bind := titleEntry(txt)
	//creates icon for notebox
	icon2 := createIcon()

	basedir, _ := getBase()
	notes, _ := read4rmBook(basedir + txt + ".txt")
	oneAdd := widget.NewButtonWithIcon("New Line", theme.ContentAddIcon(), func() {})

	// Reloaded for each new points. Holds boxbox.
	noteBox := container.NewVBox()
	//Holds note-entry widget and bullet icon
	boxBox := container.NewBorder(nil, nil, nil, nil)
	butnMap := make(map[*widget.Button]int)

	for index, item := range notes {
		delButn := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {})
		DelIcon := container.NewHBox(delButn)
		noteEntry := widget.NewEntry()
		noteEntry.SetText(item)

		boxBox = container.NewBorder(nil, nil, icon2, DelIcon, noteEntry)
		butnMap[delButn] = index
		noteBox.Add(boxBox)
		fmt.Println(butnMap)
		delButn.OnTapped = func() {
			val := butnMap[delButn]
			noteBox.Objects = removeElementByIndex(noteBox.Objects, val)
		}
	}

	// Button to create new points

	oneAdd.OnTapped = func() {
		del := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {})
		noteEntry := widget.NewEntry()
		DelIcon := container.NewHBox(del)
		boxBox = container.NewBorder(nil, nil, icon2, DelIcon, noteEntry)

		noteBox.Add(boxBox)
		del.OnTapped = func() {

		}
	}

	saveButn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		text, _ := bind.Get()
		err := titleWriter(txt, text)
		fmt.Println(err)

	})
	saveButn.Resize(fyne.NewSize(30, 30))
	DictionButn := widget.NewButton("Dictionaries", func() {})

	butnLine := container.NewBorder(nil, nil, oneAdd, DictionButn)

	fLine := container.NewVBox(title, saveButn)
	topQuater := container.NewVBox(fLine, butnLine)
	rightHand2 := container.NewBorder(topQuater, nil, nil, nil, noteBox)
	rightScroll := container.NewScroll(rightHand2)

	return rightScroll
}

func leftSide(cont *fyne.Container) fyne.CanvasObject {
	addButn := widget.NewToolbar()
	leftHand := container.NewBorder(addButn, nil, nil, nil)

	basedir, _ := getBase()
	names := dirIterator(basedir)

	l := widget.NewList(
		func() int { return len(names) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Untitled")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(names[lii])
		})
	l.OnSelected = func(id widget.ListItemID) {
		rL := rightSide(names[id], basedir)
		cont.RemoveAll()
		cont.Add(rL)

	}
	addButn.Append(widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		names = append(names, "Untitled")
		leftHand.Objects = append(leftHand.Objects)
	}))
	leftHand.Objects = append(leftHand.Objects, l)
	lScroll := container.NewScroll(leftHand)
	return lScroll
}

func loadUI() fyne.CanvasObject {
	fsttext := container.NewCenter(widget.NewLabel("Please Select a book!"))
	emptyCont := container.NewBorder(nil, nil, nil, nil, fsttext)

	l := leftSide(emptyCont)

	simp := container.NewHSplit(l, emptyCont)
	simp.Offset = 0.25

	return simp
}

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

	fullWind := container.NewBorder(header(), nil, nil, nil, loadUI())
	wind.SetContent(fullWind)
	wind.ShowAndRun()
}
