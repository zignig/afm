package main

import (
	"./fth"
	"fmt"
	"github.com/chzyer/readline"
)

func main() {
	options := fth.NewOptions("init", true, 32, 32, 4096)
	fm := fth.NewForthMachine(options)
	fm.Init()
	go fm.Run()
	go Console(fm)
	var test chan bool
	<-test
}

func Console(fm *fth.ForthMachine) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 "> ",
		HistoryFile:            "./history",
		AutoComplete:           fm.Complete,
		DisableAutoSaveHistory: true,
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
