package main

import (
	"io"
	"os"
)

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

// checkSize returns the size of a file
func checkSize(name string) (int64, error) {
	file, err := os.Stat(name)
	if err != nil {
		return 0, err
	}

	size := file.Size()
	return size, nil
}

func read4rmBook(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	if err == io.EOF {
		file.Close()
	}
	lenght, _ := checkSize(name)
	data := make([]byte, lenght)
	_, err = file.Read(data)
	if err != nil {
		return err
	}

	return nil

}

func write2Book(name, info string) error {
	//when best to use append vs write only
	opBook, err := os.OpenFile("/bookpoints/"+name+".", os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	currSize, _ := checkSize(name)
	_, wrErr := opBook.WriteAt([]byte(info), currSize+1)
	if wrErr != nil {
		return wrErr
	}
	//ioutil.WriteFile()
	return nil
}

func createBook(name string) error {
	_, err := os.Create("filepath" + name + ".") //don't forget the path and file type
	return err
}

func testData() *bookComp {
	return &bookComp{
		Comp: []*Book{
			{Title: "TestBook1"},
			{Title: "TestBook2"},
		},
	}
}

/*
When the app is opened,
*/
