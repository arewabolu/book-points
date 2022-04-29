package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//add a field of type driver

// grid button implementation
/*func addNewBooks(n int) (fyne.CanvasObject,int) {
	butt := widget.NewButtonWithIcon("add new book", theme.ContentAddIcon(), func() {

	})
	return butt
}
func bookList(books []*book) fyne.CanvasObject {
	//rect := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 255})
	//rect.Resize(fyne.NewSize(100, 100))

	size := len(books)

	//Layout of Books Notes created in the homepage
	fGrid := container.NewGridWrap(fyne.NewSize(100, 100))

	bookLayout := container.NewGridWithColumns(size, fGrid)

	//loop through all items in books
	//add to the book layout a button
	for _, n := range books {
		fGrid.Add(widget.NewButton(n.title, func() {}))
	}

	//bookList := container.NewAdaptiveGrid(size, bookLayout)
	//theme.DocumentSaveIcon()
	return bookLayout
}*/

func header() fyne.CanvasObject {
	rect := canvas.NewRectangle(color.White)
	rect.StrokeColor = color.Black
	rect.StrokeWidth = 1
	width := rect.MinSize().Width

	rect.Move(fyne.NewPos(5, 5))
	rect.Resize(fyne.NewSize(width, 90))
	//text-header for application is centerd in a rectangle

	topic := canvas.NewText("Books", color.Black)
	topic.Alignment = fyne.TextAlignCenter
	//using maxLayout to stack text on top of rectangle
	header := container.New(layout.NewMaxLayout(), rect, topic)

	return header
}

func (a *uiComp) loadUI() fyne.CanvasObject {
	var lHandle bookComp

	/*LEFTSIDE*/
	//toolbar for left-side
	leftBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			lHandle.add()
			a.bookData.Refresh()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {}),
	)

	a.bookData = widget.NewList(
		func() int { return len(lHandle.Comp) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},

		func(lii widget.ListItemID, co fyne.CanvasObject) {

			co.(*widget.Label).SetText(lHandle.Comp[lii].Title)

		},
	)

	left := container.NewBorder(leftBar, nil, nil, nil, a.bookData)

	/*RIGHT SIDE*/
	//Entry to edit title

	//svButt := widget.NewButton("Save", func(
	//implement a logger to file (maybe a json file)
	//) {})
	a.uiTitle = widget.NewEntry()
	a.uiTitle.OnChanged = func(s string) {

	}

	dol := container.NewVBox(container.NewBorder(nil, nil, nil, nil))

	//fyne.KeyReturn || fyne.Re
	//right :=
	/*JOINT SPLIT*/
	split := container.NewHSplit(left, dol)
	split.Offset = 0.25
	content := container.NewBorder(header(), nil, nil, nil, split)
	return content
}

//bullet points for chapter

func main() {
	app := app.New()
	wind := app.NewWindow("BookTakes")
	ui := &uiComp{}

	wind.Resize(fyne.NewSize(600, 600))
	wind.SetContent(ui.loadUI())
	wind.ShowAndRun()
}

/*
-User can add name of book read
-User can take one line notes from the book read
- User can list chapters where notes where taken from
-User can see a list of books
*/

/*
-App should store data in a text file?
-App should add and delete books
-New feature:saving should be a list;as pdf or just save,
	-save alone would save on device;
	-as pdf will save to a folder in device
*/
