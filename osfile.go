package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Returns the home directory of users os
func getHome() (string, error) {
	home, err := os.UserHomeDir()

	return home, err
}

// Return file directory for notes.
func getBase() (string, error) {
	home, err := getHome()
	basedir := home + "/booktakes/"
	return basedir, err
}

// Used to create Notes directory at the start of the App
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

// returns a list of names of files in basedir
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

// should be run concurrently?
// writes to info to name(bookfile)
// called when save is tapped
func write2Book(title string, noteList []string) {
	flPath, _ := getBase()
	//when best to use append vs write only
	titleTxt := flPath + title + ".txt"
	openBook, err := os.OpenFile(titleTxt, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer openBook.Close()
	for _, noteLines := range noteList {
		openBook.WriteString(noteLines + "\n")

	}
}

func getNoteList() []string {
	baseDir, _ := getBase()
	names := dirIterator(baseDir)
	return names
}

func titleWriter(oldTitle, nwTitle string) (string, error) {
	flPath, _ := getBase()
	oldTitleTxt := oldTitle + ".txt"
	nwTitleTxt := nwTitle + ".txt"

	switch {
	case nwTitle == "":
		_, err := os.Create(flPath + oldTitleTxt)
		return oldTitle, err

	case oldTitle == "":
		_, err := os.Create(flPath + nwTitleTxt)
		return nwTitle, err

	case oldTitle == "Untitled" && len(nwTitle) > 1:
		_, err := os.Create(flPath + nwTitleTxt)
		return nwTitle, err

	case nwTitle == oldTitle:
		return nwTitle, nil

	case nwTitle != oldTitle:
		err := os.Rename(flPath+oldTitleTxt, flPath+nwTitleTxt)
		return nwTitle, err

	}

	return "", nil
}

// Reads notes from files and returns
// every new line as an item in a slice
func read4rmBook(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	lineText := make([]string, 0)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		lineText = append(lineText, scanner.Text())

	}

	return lineText, nil
}

func delItem(text string) error {
	baseDir, _ := getBase()
	err := os.Remove(baseDir + text + ".txt")
	return err
}
