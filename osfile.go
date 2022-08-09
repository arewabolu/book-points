package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func writer(wr *os.File, s string) {
	_, err := wr.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
	}
}

func getHome() (string, error) {
	home, err := os.UserHomeDir()

	return home, err
}

//Return file directory for notes.
func getBase() (string, error) {
	home, err := getHome()
	basedir := home + "/booktakes/"
	return basedir, err
}

//Used to create Notes directory at the start of the App
// If directory already exists, It does nothing and returns nil.
func makeDir() error {
	//test on windows
	home, homeDirErr := os.UserHomeDir()
	if homeDirErr != nil {
		return homeDirErr
	}
	err := os.MkdirAll(home+"/booktakes/", os.ModePerm)
	return err
}

//returns a list of files in basedir
func dirIterator(basedir string) []string {
	folder, _ := os.ReadDir(basedir)
	nameSlice := make([]string, 0)
	for _, dirFile := range folder {
		if strings.HasSuffix(dirFile.Name(), ".txt") {
			name := strings.TrimSuffix(dirFile.Name(), ".txt")
			nameSlice = append(nameSlice, name)
		}
	}
	return nameSlice
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
	flPath, _ := getBase()
	nwTitle = nwTitle + ".txt"
	_, err := os.Create(flPath + nwTitle)
	if nwTitle != "" {
		oldTitle = oldTitle + ".txt"
		err := os.Rename(flPath+oldTitle, flPath+nwTitle)
		return err
	}
	return err
}

//Reads notes from files and returns
//every new line as an item in a slice
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
