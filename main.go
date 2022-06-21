package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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
func tEntry() (*widget.Entry, binding.String) {
	str := binding.NewString()

	popEntry := widget.NewEntryWithData(str)
	popEntry.Resize(fyne.NewSize(250, 30))
	go func() {
		popEntry.OnChanged = func(s string) {
			str.Set(s)
		}
	}()

	return popEntry, str
}

//Loads when note buuton is clicked
func rightSide() fyne.CanvasObject {
	icon2 := createIcon()
	titleEntry, bindStr := tEntry()
	titleEntry.Resize(fyne.NewSize(250, 30))

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
	SaveButn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		txt, _ := bindStr.Get()

		write2Titile(txt, txt)
	})
	SaveButn.Resize(fyne.NewSize(30, 30))
	DictionButn := widget.NewButton("Dictionaries", func() {})

	ButnLine := container.NewBorder(nil, nil, oneAdd, SaveButn, DictionButn)

	topQuater := container.NewVBox(titleEntry, ButnLine)
	rightHand2 := container.NewBorder(topQuater, nil, nil, nil, noteBox)

	return rightHand2
}

func leftSide(cont *fyne.Container) fyne.CanvasObject {
	lst := container.NewVBox()

	addButn := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {

			lstButn := &widget.Button{}
			lstButn.Text = "Untitiled"
			r := rightSide()
			lstButn.OnTapped = func() {

				//time.Sleep(10 * time.Millisecond)

				lstButn.Importance = widget.HighImportance
				lstButn.Refresh()

				for _, elem := range cont.Objects {
					cont.Remove(elem)
				}
				cont.Objects = append(cont.Objects, r)
			}
			lstButn.Importance = widget.LowImportance

			//for the left side
			lst.Objects = append(lst.Objects, lstButn)

		}),
	)

	leftHand := container.NewBorder(addButn, nil, nil, nil, lst)
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
	if wind.FullScreen() != true {
		wind.SetMaster()
		//wind.SetFullScreen(true)
	}
	wind.Resize(fyne.NewSize(600, 600))

	fullWind := container.NewBorder(header(), nil, nil, nil, loadUI())
	wind.SetContent(fullWind)
	wind.ShowAndRun()
}
