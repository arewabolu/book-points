package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

// checkSize returns the size of a file
func checkSize(name string) (int64, error) {
	file, err := os.Stat(name)
	if err != nil {
		return 0, err
	}

	size := file.Size()
	return size, nil
}

//should be run concurrently?
func read4rmBook(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	if err == io.EOF {
		file.Close()
	}
	defer file.Close()

	lenght, _ := checkSize(filepath)

	data := make([]byte, lenght)
	_, err = file.Read(data)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil

}

func write2Book(name, info string) error {
	//when best to use append vs write only
	openBook, err := os.OpenFile("/bookpoints/"+name+".", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	currSize, _ := checkSize(name)
	_, wrErr := openBook.WriteAt([]byte(info), currSize+1)
	if wrErr != nil {
		return wrErr
	}
	//ioutil.WriteFile()
	return nil
}

//Might nor be Needed use Write2Book.
func createBook(name string) error {
	_, err := os.Create("bootakes_" + name + ".txt") //don't forget the path and file type
	return err
}
