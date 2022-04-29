package main

import "fyne.io/fyne/v2/widget"

type uiComp struct {
	data                       *bookComp
	bookData                   *widget.List
	uiTitle, uiChapter, uiNote *widget.Entry
}
