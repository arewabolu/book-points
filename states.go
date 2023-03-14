package main

import goh "github.com/arewabolu/GoHaskell"

type bookinfo struct {
	name  string
	notes []string
}

func (b *bookinfo) remove(index int) {
	b.notes = goh.Pop(b.notes, index)
}

func (b *bookinfo) setNotes(notes []string) {
	b.notes = notes
}

func (b *bookinfo) addNotes() {
	b.notes = append(b.notes, "")
}
