package main

import (
	"bufio"
	"fmt"
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

func writer(wr *os.File, s string) {
	_, err := wr.WriteString(s)
	if err != nil {
		fmt.Println(err)
	}
}

//should be run concurrently?

func write2Book(name, info string) {
	//when best to use append vs write only

	//opens book with title to write title on
	title := name + ".txt"
	openBook, err := os.OpenFile(title, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	writer(openBook, info)

	//ioutil.WriteFile()
}

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
