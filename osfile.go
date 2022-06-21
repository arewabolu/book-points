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

func write2Titile(name, info string) {
	//when best to use append vs write only
	title := name + "\n" + ".txt"
	openBook, err := os.OpenFile(title, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	//currSize, _ := checkSize(name)
	_, wrErr := openBook.Write([]byte(info))
	if wrErr != nil {
		fmt.Println(wrErr)
	}
	//ioutil.WriteFile()
}
