package main

import (
	"../afm"
	"fmt"
	"github.com/chzyer/readline"
	"io"
)

func main() {
	exit := make(chan bool, 1)
	options := afm.DefaultOptions()
	fm := afm.NewForthMachine(options)
	fm.Init()
	go fm.Run(exit)
	c := NewConsole(fm, exit)
	go c.Run()
	<-exit
	c.Close()
	fmt.Println("EXITING")
}

type Console struct {
	rl   *readline.Instance
	exit chan bool
	fm   *afm.ForthMachine
}

func NewConsole(fm *afm.ForthMachine, exit chan bool) (c *Console) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 fm.Options.Prompt,
		HistoryFile:            "./history",
		AutoComplete:           fm.Complete,
		DisableAutoSaveHistory: true,
	})
	if err != nil {
		panic(err)

	}
	c = &Console{
		rl:   rl,
		exit: exit,
		fm:   fm,
	}
	return c
}

func (c *Console) Close() {
	c.rl.Close()
}

func (c *Console) Run() {
	rl := c.rl
	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println(err)
			if err == readline.ErrInterrupt {
				fmt.Println("INTERRUPT")
			}
			if err == io.EOF {
				fmt.Println("EXIT")
				c.exit <- true
				break
			}
		}
		rl.SaveHistory(line)
		c.fm.Input <- line
		if c.fm.Exit {
			fmt.Println("console exit")
			rl.Close()
			c.exit <- true
			break
		}
	}
	fmt.Println("EXIT CONSOLE")
}
