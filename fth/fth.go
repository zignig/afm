package fth

import (
	"fmt"
)

func (fm *ForthMachine) Run(exit chan bool) (e error) {
	fm.Words()
	fmt.Printf("Starting on %s\n", fm.startword)
	w, err := fm.d.Search(fm.startword)
	if err != nil {
		fmt.Println("boot error :", err)
		return err
	}
	w.Do()
	err = fm.LoadFile("./base.fth")
	if err != nil {
		return err
	}
	for {
		if fm.exit {
			exit <- true
			break
		}
		select {
		case line := <-fm.Input:
			fmt.Println(line)
			fm.GetLine(line)
			fm.Process()
		}
		fmt.Println("ok")
	}
	fmt.Println("EXIT MACHINE")
	return
}
