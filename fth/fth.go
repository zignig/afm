package fth

import (
	"fmt"
)

func (fm *ForthMachine) Run(exit chan bool) (e error) {
	err := fm.LoadFile("./base.fth")
	if err != nil {
		return err
	}
	fmt.Printf("Starting on %s\n", fm.startword)
	w, err := fm.d.Search(fm.startword)
	if err != nil {
		fmt.Println("boot error :", err)
		return err
	}
	w.Do()
	for {
		if fm.Exit {
			exit <- true
			break
		}
		select {
		case line := <-fm.Input:
			fm.GetLine(line)
			fm.Process()
		}
	}
	fmt.Println("EXIT MACHINE")
	return
}
