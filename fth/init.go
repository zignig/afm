package fth

import (
	"fmt"
)

// initialization things

func (fm *ForthMachine) Init() {
	init := NewBaseWord("init")
	initTest := func() (e error) {
		fmt.Println("this is a inderected wooty function")
		return
	}
	init.SetExec(initTest)
	fm.Add(init)

	words := NewBaseWord("words")
	fm.Add(words)

	wd := func() (e error) {
		fm.Words()
		return
	}
	words.SetExec(wd)

	help := NewBaseWord("help")
	fm.Add(help)
}
