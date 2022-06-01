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

var (
	Title string
	size  int
)

//type Num struct {
//	num int
//}

//func (n *Num) incr() {
//	n.num++
//}

func incr2(p *int) {
	size++
	*p = size
}

func chngetTitle(s *string) {
	*s = Title
}
func setDefault() string {
	return "Untitled"
}

func createIcon() *widget.Icon {
	resc, _ := LoadResourceFromPath("/home/arthemis/Pictures/icons/bullet")
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
func loadUI() fyne.CanvasObject {

	lst := container.NewVBox()
	addButn := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			listbtn := widget.NewButton(setDefault(), func() {})
			listbtn.Alignment = widget.ButtonAlignLeading
			lst.Objects = append(lst.Objects, listbtn)
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

	//Create New Entry/Input widget
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter your name")
	entry.Resize(fyne.NewSize(250, 30))

	noteBox := container.NewVBox()
	boxBox := container.NewBorder(nil, nil, nil, nil)
	noteTlBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			icon2 := createIcon()
			noteEntry := widget.NewEntry()
			boxBox = container.NewBorder(nil, nil, icon2, nil, noteEntry)

			noteBox.Add(boxBox)
		}),

		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), func() {
			icon2 := createIcon()
			lNoteEntry := widget.NewMultiLineEntry()
			boxBox = container.NewBorder(nil, nil, icon2, nil, lNoteEntry)
			noteBox.Add(boxBox)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {}),
	)
	//nwntButt.Alignment = widget.ButtonAlignTrailing

	topQuater := container.NewVBox(entry, noteTlBar)
	rightHand := container.NewBorder(topQuater, nil, nil, nil, noteBox)

	simp := container.NewHSplit(leftHand, rightHand)
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
