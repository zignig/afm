package afm

import (
	"fmt"
)

func (fm *ForthMachine) Run(exit chan bool) (e error) {
	err := fm.LoadFile("./base.fth")
	if err != nil {
		fm.out(err)
	}
	fm.outf("Starting on %s\n", fm.startword)
	w, err := fm.d.Search(fm.startword)
	if err != nil {
		fm.out("boot error :", err)
		return err
	}
	w.Do() // execute the init
	for {
		if fm.Exit {
			exit <- true
			break
		}
		select {
		case line := <-fm.Input:
			fm.GetLine(line)
			err = fm.Process()
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	fm.out("EXIT MACHINE")
	return
}

func (fm *ForthMachine) Process() (err error) {
	for {
		tok, empty := fm.NextToken()
		if empty {
			return
		}
		w, err := fm.d.Search(tok)
		if debug {
			fm.out(w, err, tok)
		}
		// TODO check for numbers
		if err != nil {
			fm.out(tok, "--", err)
			fm.compile = false
			return err
		} else {
			if fm.compile {
				// compiler
				if w.IsImm() {
					if debug {
						fm.out("imm function ", w.Name())
					}
					err = w.Do()
					if err != nil {
						return err
					}
				}
				if debug {
					fm.out("compile", fm.current, tok, w)
				}
				if fm.current != nil {
					fm.current.Add(w)
				}
			} else {
				// interpreter
				// set the program counter
				fm.pc = &PCRef{w: w, offset: 0}
				// execute the woord
				err = w.Do()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (fm *ForthMachine) Call() {
	fmt.Println("Calling function for primary execution")
	call := func() (e error) {
		//fmt.Println("push onto return stack")
		fm.rStack.Push(fm.pc)
		fm.rStack.Show()
		fmt.Print("RUN > ")
		//push the literal on the stack
		if fm.pc.w.IsLit() {
			e = fm.dStack.Push(fm.pc.w)
			if e != nil {
				return e
			}
		}
		// execute
		// run through and call words here
		// will recurse down and back and up
		fmt.Println(fm.pc.w)
		fm.pc.inc()
		// need to increment the PC counter with bounds checking
		//fmt.Println("pop return stack")

		newPc, err := fm.rStack.Pop()
		if err != nil {
			return err
		}
		//fmt.Println("update program counter")
		fm.pc = newPc
		fm.rStack.Show()
		return
	}
	// assign this function to the global call
	// it is attached to any normal defined word
	fm.call = call
}
