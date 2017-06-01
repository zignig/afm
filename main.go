package main

import (
	"./fth"
	"fmt"
	"github.com/chzyer/readline"
)

func main() {
	exit := make(chan bool, 1)
	options := afm.DefaultOptions()
	fm := afm.NewForthMachine(options)
	fm.Init()
	go fm.Run(exit)
	go Console(fm, exit)
	<-exit
	fmt.Println("EXITING")
}
func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
		//	case readline.CharInterrupt:
		////		return r, false
	}
	return r, true
}

func Console(fm *afm.ForthMachine, exit chan bool) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 fm.Options.Prompt,
		HistoryFile:            "./history",
		AutoComplete:           fm.Complete,
		DisableAutoSaveHistory: true,
		FuncFilterInputRune:    filterInput,
	})
	if err != nil {
		panic(err)

	}
	defer rl.Close()
	for {
		if fm.Exit {
			fmt.Println("console exit")
			//			exit <- true
			return
		}
		line, err := rl.Readline()
		if err != nil {
			return
		}
		rl.SaveHistory(line)
		fm.Input <- line
	}
}
