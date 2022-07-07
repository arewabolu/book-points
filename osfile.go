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

func getBase() string {
	home, _ := os.UserHomeDir()
	basedir := home + "/booktakes/"
	return basedir
}

func makeDir() error {
	//test on windows
	home, _ := os.UserHomeDir()
	err := os.MkdirAll(home+"/booktakes/", os.ModePerm)

	return err
}

//should be run concurrently?
//writes to info to name(bookfile)
//called when save is clicked
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

func titleWriter(oldTitle, nwTitle string) error {
	flPath := getBase()
	oldTitle = oldTitle + ".txt"
	nwTitle = nwTitle + ".txt"
	err := os.Rename(flPath+oldTitle, flPath+nwTitle)
	return err
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
