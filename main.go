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

type book struct {
	title   string
	chapter int
	lesson
}

type lesson struct {
	notes string
}

func addNewBooks(n int) (fyne.CanvasObject, int) {
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
}

func points() fyne.CanvasObject {
	//newPoint := widget.NewIcon(theme.MailSendIcon())
	content := widget.NewEntry()
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.AccountIcon(), func() {
			content.SetText("")
		}))
	return container.New(layout.NewBorderLayout(bar, content, nil, nil), bar, content)
}

func appTabButton() fyne.CanvasObject {
	tab := container.NewAppTabs(container.NewTabItemWithIcon("book", theme.CancelIcon(), points()))
	//tab.CreateRenderer()
	tab.SetTabLocation(container.TabLocationTop)
	return tab
}

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

func main() {
	app := app.New()
	wind := app.NewWindow("BookTakes")

	list := []*book{
		{title: "book1", chapter: 1},
		{title: "book2", chapter: 2, lesson: lesson{notes: ""}},
		{title: "book3", chapter: 1},
		{title: "book4", chapter: 2, lesson: lesson{notes: ""}},
	}

	content := container.NewBorder(header(), homeAddButton(), nil, nil, bookList(list))
	wind.Resize(fyne.NewSize(600, 600))
	wind.SetContent(content)
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
*/
