package main

import (
	"../afm"
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

func Console(fm *afm.ForthMachine, exit chan bool) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 fm.Options.Prompt,
		HistoryFile:            "./history",
		AutoComplete:           fm.Complete,
		DisableAutoSaveHistory: true,
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
