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
	_, err = fm.d.Search(fm.startword)
	if err != nil {
		fm.out("boot error :", err)
		return err
	}
	// this is the main machine loop
	for {
		if fm.Exit {
			exit <- true
			break
		}

		select {
		case line := <-fm.Input: // a string of input has arrived
			fm.GetLine(line)
			err = fm.Process()
			// more input sources here
			// io
			// nonvolatile memory
			// random go struct with heaps of goroutines running ?
		// main execution loop goes here
		case xt := <-fm.XT: // grab an execution token out
			err = fm.Exec(xt)
		}

		if err != nil {
			// need error handlers here (perhaps dynamic)
			fmt.Println(err)
		}
		fmt.Println("ok")
	}
	fm.out("EXIT MACHINE")
	return
}

func (fm *ForthMachine) Exec(w Word) (err error) {
	fmt.Print("EXECUTE THIS >")
	err = w.Do()
	fmt.Println(w, err)
	return err
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
			// always execute immediate words
			if w.IsImm() {
				if debug {
					fm.out("imm function ", w.Name())
				}
				err = fm.Exec(w)
				if err != nil {
					return err
				}
				continue
			}
			if fm.compile {
				// compiler
				if debug {
					fm.out("compile", fm.current, tok, w)
				}
				if fm.current != nil {
					fm.current.Add(w)
				}
			} else {
				// spool into execution
				fmt.Println("Spool >", w)
				fm.XT <- w
			}
		}
	}
	return nil
}

// this structure is wrong , need to call from A hight level loop
// disabled for now
func (fm *ForthMachine) Call() {
	// Calling function for primary execution
	call := func() (e error) {
		//fmt.Println("push onto return stack")
		e = fm.rStack.Push(fm.pc.wrap())
		if e != nil {
			return e
		}
		// show the rstack
		fm.rStack.Show()
		fmt.Print("RUN > ")
		//push the literal on the stack
		current := fm.pc.w
		if current.IsLit() {
			e = fm.dStack.Push(current)
			if e != nil {
				return e
			}
		}
		// execute
		// move the pragram counter to the start of the next word
		fm.pc.Set(current, 0)
		// execute
		e = fm.pc.w.Do()
		if e != nil {
			return e
		}
		// run through and call words here
		// will recurse down and back and up
		fmt.Println(fm.pc)
		// need to increment the PC counter with bounds checking
		e = fm.pc.inc()
		if e != nil {
			return e
		}
		//show the rstack again
		fm.rStack.Show()
		return
	}
	// assign this function to the global call
	// it is attached to any normal defined word
	fm.call = call
}
