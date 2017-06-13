package afm

import (
	"fmt"
)

// Run is the main forth machine loop
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
			fm.pc.Set(xt, 0)
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
	if w.Length() > 0 {
		err = fm.Composite(w)
	} else {
		fmt.Print("EXEC>")
		err = w.Do()
	}
	return err
}

// Execute a composite word
func (fm *ForthMachine) Composite(w Word) (err error) {
	for i := 0; i < w.Length(); i++ {
		nw, err := w.Get(i)
		if err != nil {
			return err
		}
		fmt.Println("run >", nw)
	}
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

func (fm *ForthMachine) Call() {
	// Calling function for primary execution
	call := func() (e error) {
		fmt.Println("CALL")
		return
	}
	fm.call = call
}
