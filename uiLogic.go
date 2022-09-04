package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

// text-header for application
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

func titleEntry(text string) (*widget.Entry, binding.String) {
	titleEntry := widget.NewEntry()
	titleBind := binding.NewString()
	titleEntry.SetText(text)
	titleEntry.Resize(fyne.NewSize(250, 30))
	titleEntry.OnChanged = func(s string) {
		titleBind.Set(s)
		titleEntry.Bind(titleBind)
	}
	return titleEntry, titleBind
}

// Loads when note buuton is clicked
func loadRightSide(name []string, ID int, w fyne.Window) fyne.CanvasObject {
	baseDir, _ := getBase()
	title, titleBind := titleEntry(name[ID])
	icon2 := createIcon()
	boxBox := container.NewBorder(nil, nil, nil, nil)
	notes, _ := read4rmBook(baseDir + name[ID] + ".txt")
	noteBindings := binding.BindStringList(&notes)
	oneAdd := widget.NewButtonWithIcon("New Line", theme.ContentAddIcon(), func() {
		noteBindings.Append("")
	})

	//A trick for managing each lines id

	listing := widget.NewListWithData(
		noteBindings,
		func() fyne.CanvasObject {
			noteEntry := widget.NewEntry()
			delButn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				position := slices.Index(notes, noteEntry.Text)
				nwNote := removeElementByIndex(notes, position)
				noteBindings.Set(nwNote)
			})
			delButn.Resize(fyne.NewSize(10, 10))

			boxBox = container.NewBorder(nil, nil, icon2, delButn, noteEntry)
			return boxBox
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*fyne.Container).Objects[0].(*widget.Entry).Bind(di.(binding.String))
		},
	)

	saveButn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		saveFunc(name[ID], titleBind, noteBindings, w)

	})
	saveButn.Resize(fyne.NewSize(30, 30))

	DictionButn := widget.NewButton("Dictionaries", func() {})
	butnLine := container.NewBorder(nil, nil, oneAdd, DictionButn)

	fLine := container.NewVBox(title, saveButn)
	topQuater := container.NewVBox(fLine, butnLine)
	rightHand2 := container.NewBorder(topQuater, nil, nil, nil, listing)
	rightScroll := container.NewScroll(rightHand2)

	return rightScroll
}

func leftSide(cont *fyne.Container, w fyne.Window) fyne.CanvasObject {

	nameList := widget.NewListWithData(
		bindingFunc(),
		func() fyne.CanvasObject {
			label := widget.NewLabel("Untitled")
			noteDel := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				delItem(label.Text)

			})
			noteDel.Resize(fyne.NewSize(10, 10))

			return container.NewBorder(nil, nil, label, noteDel)
		},
		// binds the template item above to a string binding
		func(data binding.DataItem, co fyne.CanvasObject) {
			co.(*fyne.Container).Objects[0].(*widget.Label).Bind(data.(binding.String))
		},
	)
	nameList.OnSelected = func(id widget.ListItemID) {
		cont.RemoveAll()
		cont.Add(loadRightSide(getNoteList(), id, w))
	}
	nameList.OnUnselected = func(id widget.ListItemID) {
		cont.RemoveAll()
	}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			//time.NewTicker(5 * time.Second)
			bindingFunc().Reload()
		}
	}()

	addButn := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {

			bindingFunc().Append("Untitled")
		}))
	leftHand := container.NewBorder(addButn, nil, nil, nil, nameList)
	lScroll := container.NewScroll(leftHand)
	return lScroll
}

func loadUI(w fyne.Window) fyne.CanvasObject {
	fsttext := container.NewCenter(widget.NewLabel("Please Select a book!"))
	emptyCont := container.NewBorder(nil, nil, nil, nil, fsttext)

	l := leftSide(emptyCont, w)

	simp := container.NewHSplit(l, emptyCont)
	simp.Offset = 0.25

	return simp
}
