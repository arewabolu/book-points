package main

import (
	"errors"
	"fmt"
	"image/color"
	"os"
	"strings"

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

func rightSidewithFile(txt string) fyne.CanvasObject {
	icon2 := createIcon()
	titleEntry := widget.NewEntry()
	titleBind := binding.NewString()
	titleEntry.Resize(fyne.NewSize(250, 30))
	titleEntry.Text = txt
	titleEntry.OnChanged = func(s string) {
		titleBind.Set(s)
		titleEntry.Bind(titleBind)
	}
	titleEntry.Refresh()

	// Reloaded for each new points
	noteBox := container.NewVBox()
	//Holds note-entry widget and bullet icon
	boxBox := container.NewBorder(nil, nil, nil, nil)
	// Button to create new points
	oneAdd := widget.NewButtonWithIcon("New Line", theme.ContentAddIcon(), func() {
		del := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {})
		noteEntry := widget.NewEntry()
		noteEntry.OnChanged = func(s string) {

		}
		DelIcon := container.NewHBox(del)
		boxBox = container.NewBorder(nil, nil, icon2, DelIcon, noteEntry)
		noteBox.Add(boxBox)
	})

	basedir := getBase()
	notes, _ := read4rmBook(basedir + txt + ".txt")

	for _, item := range notes {
		noteEntry := widget.NewEntry()
		noteEntry.SetText(item)
		os.Getwd()

		del := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {})
		DelIcon := container.NewHBox(del)
		boxBox = container.NewBorder(nil, nil, icon2, DelIcon, noteEntry)
		noteBox.Add(boxBox)
		noteBox.Refresh()
	}

	saveButn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		text, _ := titleBind.Get()

		err := titleWriter(txt, text)
		fmt.Println(err)
	})
	saveButn.Resize(fyne.NewSize(30, 30))
	DictionButn := widget.NewButton("Dictionaries", func() {})

	butnLine := container.NewBorder(nil, nil, oneAdd, DictionButn)

	fLine := container.NewVBox(titleEntry, saveButn)
	topQuater := container.NewVBox(fLine, butnLine)
	rightHandOldFl := container.NewBorder(topQuater, nil, nil, nil, noteBox)
	rightScroll := container.NewScroll(rightHandOldFl)

	return rightScroll
}

//Loads when note buuton is clicked
func rightSide(txt string) fyne.CanvasObject {
	icon2 := createIcon()
	titleEntry := widget.NewEntry()
	titleBind := binding.NewString()
	titleEntry.Text = txt
	titleEntry.OnChanged = func(s string) {
		titleBind.Set(s)
		titleEntry.Bind(titleBind)

	}
	titleEntry.Resize(fyne.NewSize(250, 30))
	titleEntry.Refresh()

	// Reloaded for each new points
	noteBox := container.NewVBox()
	//Holds note-entry widget and bullet icon
	boxBox := container.NewBorder(nil, nil, nil, nil)
	// Button to create new points
	oneAdd := widget.NewButtonWithIcon("New Line", theme.ContentAddIcon(), func() {
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

	saveButn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {

	})
	saveButn.Resize(fyne.NewSize(30, 30))
	DictionButn := widget.NewButton("Dictionaries", func() {})

	butnLine := container.NewBorder(nil, nil, oneAdd, DictionButn)

	fLine := container.NewVBox(titleEntry, saveButn)
	topQuater := container.NewVBox(fLine, butnLine)
	rightHand2 := container.NewBorder(topQuater, nil, nil, nil, noteBox)

	return rightHand2
}

func leftSide(cont *fyne.Container) fyne.CanvasObject {

	lst := container.NewVBox()
	basedir := getBase()
	folder, _ := os.ReadDir(basedir)
	for _, dirFile := range folder {
		if strings.HasSuffix(dirFile.Name(), ".txt") {
			name := strings.TrimSuffix(dirFile.Name(), ".txt")
			r := rightSidewithFile(name)
			oldFlButn := widget.NewButton(name, func() {

				for _, elem := range cont.Objects {
					cont.Remove(elem)
				}
				cont.Add(r)
			})
			oldFlButn.Refresh()
			lst.Objects = append(lst.Objects, oldFlButn)

		}

	}

	addButn := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {

			lstButn := &widget.Button{}
			lstButn.Text = "Untitled"
			r := rightSide(lstButn.Text)
			lstButn.OnTapped = func() {
				for _, elem := range cont.Objects {
					cont.Remove(elem)
				}
				cont.Objects = append(cont.Objects, r)
			}

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
