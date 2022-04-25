package main

type Book struct {
	Title   string
	Note    []string
	Chapter int
}

type bookComp struct {
	Comp []*Book
}

// used to set default title
func (b *Book) setDefTitle() {
	b.Title = "Untitled"
}

//func (b *Book) SetPoints(tittle string) string {
//	b.title = tittle
//	return b.title
//}

//append to the initial list a new Book
func (n *bookComp) add() {
	var b *Book
	b.setDefTitle()

	n.Comp = append(n.Comp, b)
}

func testData() *bookComp {
	return &bookComp{
		Comp: []*Book{
			{Title: "TestBook1"},
			{Title: "TestBook2"},
		},
	}
}
