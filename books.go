package main

type book struct {
	Title string
}

// represents a composition of the books
type books struct {
	Comp []books
}

// Creates new book
func setBooks() (b book) {
	return b
}

// for book
func (b *book) SetTitle(s string) {
	s = b.Title
}

//func Creates() *Books {
//	nw := newBook()
//	nw.Title = "Untitled"
//	nw.Note = append(nw.Note, "Type Note here")
//	return &nw
//}
