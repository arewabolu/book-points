package main

import (
	"errors"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
func loadRightSide(txt, dir string, ID int) fyne.CanvasObject {
	title, titleBind := titleEntry(txt)
	//creates icon for notebox
	icon2 := createIcon()

	basedir, _ := getBase()
	notes, _ := read4rmBook(basedir + txt + ".txt")
	oneAdd := widget.NewButtonWithIcon("New Line", theme.ContentAddIcon(), func() {})

	//noteBox := container.NewVBox()

	boxBox := container.NewBorder(nil, nil, nil, nil)
	noteBindings := binding.BindStringList(&notes)

	//A trick for managing each lines id

	listing := widget.NewListWithData(
		noteBindings,
		func() fyne.CanvasObject {
			delButn := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {})
			delIcon := container.NewHBox(delButn)
			noteEntry := widget.NewEntry()

			boxBox = container.NewBorder(nil, nil, icon2, delIcon, noteEntry)
			return boxBox
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*fyne.Container).Objects[0].(*widget.Entry).Bind(di.(binding.String))
		},
	)

	// Button to create new points

	oneAdd.OnTapped = func() {
		noteBindings.Append("")

	}

	saveButn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		titleText, _ := titleBind.Get()
		writtenTitle, titleWriterErr := titleWriter(txt, titleText)
		fmt.Println(titleWriterErr)
		noteList, err2 := noteBindings.Get()
		fmt.Println(err2)
		write2Book(writtenTitle, noteList)

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

func leftSide(cont *fyne.Container) fyne.CanvasObject {

	baseDir, _ := getBase()
	names := dirIterator(baseDir)

	nameBinding := binding.BindStringList(&names)

	nameList := widget.NewListWithData(
		nameBinding,
		func() fyne.CanvasObject {
			return widget.NewLabel("Untitled")
		},
		// binds the template item above to a string binding
		func(data binding.DataItem, co fyne.CanvasObject) {
			co.(*widget.Label).Bind(data.(binding.String))
		},
	)
	nameList.OnSelected = func(id widget.ListItemID) {
		rL := loadRightSide(names[id], baseDir, id)
		cont.RemoveAll()
		cont.Add(rL)
	}
	nameList.OnUnselected = func(id widget.ListItemID) {
		cont.RemoveAll()
	}

	addButn := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {

			nameBinding.Append("Untitled")
		}))
	leftHand := container.NewBorder(addButn, nil, nil, nil, nameList)
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
