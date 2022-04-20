package main

type Book struct {
	Title   string
	Note    []string
	Chapter int
}

type bookComp struct {
	Comp []*Book
}

func (b Book) SetTitle(title string) string {
	b.Title = title
	if b.Title == "" {
		return "Untitled"
	}
	return b.Title
}

//func (b *Book) SetPoints(tittle string) string {
//	b.title = tittle
//	return b.title
//}

//append to the initial list a new Book
func (n *bookComp) add(b *Book) {
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
