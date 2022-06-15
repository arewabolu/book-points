package main

import (
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

func createIcon() *widget.Icon {
	resc, _ := LoadResourceFromPath("https://github.com/arewabolu/book-points/blob/main/Assets/icons/bullet?raw=true")
	theme.NewThemedResource(resc)
	icon := widget.NewIcon(resc)

	return icon
}

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

// Creates a Title entry
func tEntry(b *book) (*widget.Entry, *book) {
	popEntry := widget.NewEntry()

	popEntry.OnChanged = func(s string) {
		SetTitle(s, b)
		popEntry.Refresh()
	}
	return popEntry, b
}

//Loads when note buuton is clicked
func rightSide(p *widget.Entry) fyne.CanvasObject {
	icon2 := createIcon()

	// Reloaded for each new points
	noteBox := container.NewVBox()
	//Holds note-entry widget and bullet icon
	boxBox := container.NewBorder(nil, nil, nil, nil)
	// Button to create new points
	oneAdd := widget.NewButton("New Line", func() {
		del := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {})
		noteEntry := widget.NewEntry()
		DelIcon := container.NewHBox(del)
		boxBox = container.NewBorder(nil, nil, icon2, DelIcon, noteEntry)

		noteBox.Add(boxBox)
		del.OnTapped = func() {
			for ind := range noteBox.Objects {
				noteBox.Remove(noteBox.Objects[ind])
			}

		}
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

	ButnLine := container.NewBorder(nil, nil, oneAdd, SaveButn, DictionButn)
	p.Resize(fyne.NewSize(250, 30))

	topQuater := container.NewVBox(p, ButnLine)
	rightHand2 := container.NewBorder(topQuater, nil, nil, nil, noteBox)

	return rightHand2
}

func loadUI(w fyne.Window) fyne.CanvasObject {
	fsttext := container.NewCenter(widget.NewLabel("Please Select a book!"))
	emptyCont := container.NewBorder(nil, nil, nil, nil, fsttext)
	lst := container.NewVBox()
	addButn := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			b := &book{}
			popEnt, bK := tEntry(b)

			fmt.Println(bK)

			NewDialog := dialog.NewForm("Create New Book", "Create", "Cancel", []*widget.FormItem{widget.NewFormItem("Name", popEnt)}, func(b bool) {
				if b == true {

				}
			}, w)
			NewDialog.Show()
			r := rightSide(popEnt)
			lstButn := widget.NewButton(b.Title, func() {

				// to replace previous right side widget
				for _, elem := range emptyCont.Objects {
					emptyCont.Remove(elem)
				}
				emptyCont.Objects = append(emptyCont.Objects, r)
			})

			//for the left side
			lst.Objects = append(lst.Objects, lstButn)
			lstButn.Refresh()
		}),
	)
	addButn.Refresh()

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
	if wind.FullScreen() != true {
		wind.SetMaster()
		wind.SetFullScreen(true)
	}
	//	wind.Resize(fyne.NewSize(600, 600))
	fullWind := container.NewBorder(header(), nil, nil, nil, loadUI(wind))
	wind.SetContent(fullWind)
	wind.ShowAndRun()
}
