package fth

import (
	"fmt"
)

// initialization things
var helpText string = " this is a the start of a abstract forth machine"

func (fm *ForthMachine) Init() {
	// the usual init work
	init := NewBaseWord("init")
	initTest := func() (e error) {
		fmt.Println("init would go here")
		return
	}
	init.SetExec(initTest)
	fm.Add(init)

	// list of current words
	words := NewBaseWord("words")
	fm.Add(words)

	wd := func() (e error) {
		fm.Words()
		return
	}
	words.SetExec(wd)

	// help function
	help := NewBaseWord("help")
	fm.Add(help)
	helpFunc := func() (e error) {
		fmt.Println(helpText)
		return
	}
	help.SetExec(helpFunc)

	// debug function
	debug := NewBaseWord("debug")
	fm.Add(debug)
	debugFunc := func() (e error) {
		fm.debug = !fm.debug
		if fm.debug {
			fmt.Println("debug on")
		} else {
			fmt.Println("debug off")
		}
		return
	}

	debug.SetExec(debugFunc)

	// define new word
	define := NewBaseWord(":")
	fm.Add(define)
	endDefine := NewBaseWord(";")
	fm.Add(endDefine)

	// bye
	bye := NewBaseWord("bye")
	fm.Add(bye)
	byeFunc := func() (e error) {
		fm.exit = true
		return
	}
	bye.SetExec(byeFunc)
}
