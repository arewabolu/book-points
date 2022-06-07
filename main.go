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

func createIcon() *widget.Icon {
	resc, _ := LoadResourceFromPath("github.com/arewabolu/book-points/blob/main/Assets/icons/bullet")
	theme.NewThemedResource(resc)
	icon := widget.NewIcon(resc)

	return icon
}

//add a field of type driver

//text-header for application
func header() fyne.CanvasObject {
	rect := canvas.NewRectangle(color.White)
	rect.StrokeColor = color.Black
	rect.StrokeWidth = 1
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
func rightSide() (fyne.CanvasObject, book) {

	// Reloaded for each new points
	noteBox := container.NewVBox()

	//Holds entry widget and bullet icon
	boxBox := container.NewBorder(nil, nil, nil, nil)

	oneAdd := widget.NewButton("New Line", func() {
		icon2 := createIcon()
		noteEntry := widget.NewEntry()
		boxBox = container.NewBorder(nil, nil, icon2, nil, noteEntry)

		noteBox.Add(boxBox)
	})

	//doubleAdd := widget.NewButton("Double", func() {
	//	icon2 := createIcon()
	//	noteEntry := widget.NewMultiLineEntry()
	//	boxBox = container.NewBorder(nil, nil, icon2, nil, noteEntry)
	//	noteBox.Add(boxBox)
	//})
	SaveButn := widget.NewButton("Save", func() {})
	SaveButn.Resize(fyne.NewSize(30, 30))
	DictionButn := widget.NewButton("Dictionaries", func() {})

	//Title Input widget
	entry := widget.NewEntry()
	entry.Resize(fyne.NewSize(250, 30))

	//Change title of Book through entry
	b := setBooks()
	entry.OnChanged = func(s string) {

		b.Title = s
	}
	ButnLine := container.NewBorder(nil, nil, oneAdd, SaveButn, DictionButn)

	topQuater := container.NewVBox(entry, ButnLine)
	rightHand2 := container.NewBorder(topQuater, nil, nil, nil, noteBox)
	return rightHand2, b
}

func loadUI() fyne.CanvasObject {
	fsttext := container.NewCenter(widget.NewLabel("Please Select a book!"))
	emptyCont := container.NewBorder(nil, nil, nil, nil, fsttext)

	lst := container.NewVBox()

	addButn := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			r, bT := rightSide()
			lstButn := widget.NewButton(bT.Title, func() {
				// to replace previous right sidr widget
				if len(emptyCont.Objects) > 0 {
					for _, elem := range emptyCont.Objects {
						emptyCont.Remove(elem)
					}
				}

				emptyCont.Objects = append(emptyCont.Objects, r)

			})
			//for the left side
			lst.Objects = append(lst.Objects, lstButn)

		}),
	)

	leftHand := container.NewBorder(
		addButn,
		nil,
		nil,
		nil,
		lst,
	)
	leftHand.Refresh()

	simp := container.NewHSplit(leftHand, emptyCont)
	simp.Offset = 0.25

	return simp
}

//bullet points for chapter

func main() {
	app := app.New()
	wind := app.NewWindow("BookTakes")
	wind.Resize(fyne.NewSize(600, 600))
	fullWind := container.NewBorder(header(), nil, nil, nil, loadUI())
	wind.SetContent(fullWind)
	wind.ShowAndRun()
}
