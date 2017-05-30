package fth

import (
	"fmt"
)

// initialization things
var helpText string = " this is a the start of a abstract forth machine"

func (fm *ForthMachine) Init() {
	// base colon and semi colon defs
	fm.SetDef()
	// extra functions
	fm.Extra()
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

	// dump code
	dump := NewBaseWord("dump")
	fm.Add(dump)
	dmp := func() (e error) {
		fm.d.dump()
		return
	}
	dump.SetExec(dmp)

	// help function
	help := NewBaseWord("help")
	fm.Add(help)
	helpFunc := func() (e error) {
		fmt.Println(helpText)
		return
	}
	help.SetExec(helpFunc)

	// debug function
	debugf := NewBaseWord("debug")
	fm.Add(debugf)
	debugFunc := func() (e error) {
		debug = !debug
		if debug {
			fmt.Println("debug on")
		} else {
			fmt.Println("debug off")
		}
		return
	}

	debugf.SetExec(debugFunc)

	// bye
	bye := NewBaseWord("bye")
	fm.Add(bye)
	byeFunc := func() (e error) {
		fmt.Println("EXIT SET")
		fm.exit = true
		return
	}
	bye.SetExec(byeFunc)
}
