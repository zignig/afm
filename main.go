package main

import (
	"./fth"
	"fmt"
	"github.com/chzyer/readline"
)

func main() {
	var exit chan bool
	options := fth.NewOptions("init", true, 32, 32, 4096)
	fm := fth.NewForthMachine(options)
	fm.Init()
	go fm.Run(exit)
	go Console(fm)
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

func Console(fm *fth.ForthMachine) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 "> ",
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
		line, err := rl.Readline()
		if err != nil {
			return
		}
		fmt.Println(line)
		rl.SaveHistory(line)
		fm.Input <- line
	}
}
