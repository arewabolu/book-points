package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// checkSize returns the size of a file

func writer(wr *os.File, s string) {
	_, err := wr.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
	}
}

//should be run concurrently?

func write2Book(name, info string) {
	//when best to use append vs write only
	title := name + ".txt"
	openBook, err := os.OpenFile(title, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer openBook.Close()
	if err != nil {
		fmt.Println(err)
	}
	writer(openBook, info)
}

func titleWriter(nwTitle string) {
	flPath := "./bookpoints/bookTitleList.txt"
	openFile, err := os.OpenFile(flPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	read4Title()
	defer openFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	writer(openFile, nwTitle)
}

func read4Title() []string {
	flPath := "./bookpoints/bookTitleList.txt"
	file, err := os.Open(flPath)
	if err != nil {
		log.Panic(err)
	}

	nwScan := bufio.NewScanner(file)
	fileSlice := make([]string, 0)

	nwScan.Split(bufio.ScanLines)
	for nwScan.Scan() {
		fileSlice = append(fileSlice, nwScan.Text())
	}
	return fileSlice
}

func read4rmBook(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	lineText := make([]string, 0)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lineText = append(lineText, scanner.Text())
	}

	return lineText, nil
}
