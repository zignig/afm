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
			break
		}

		select {
		case line := <-fm.Input: // a string of input has arrived
			fm.GetLine(line)
			err = fm.Process()
			if err != nil {
				fmt.Println("PROC ERROR", err)
			}
		// more input sources here
		// io
		// nonvolatile memory
		// random go struct with heaps of goroutines running ?
		// main execution loop goes here
		case xt := <-fm.XT: // grab an execution token out
			fm.pc.Set(xt, 0)
			err = fm.Exec(xt)
		case <-exit:
			exit <- true
			fm.Exit = true
			break
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

// calling composite word
func (fm *ForthMachine) Call() {
	// Calling function for primary execution
	call := func() (e error) {
        for {
		fmt.Println("CALL")
        fmt.Println(fm.pc)
        currentWord , e  := fm.pc.Get()
        fmt.Println(e,fm.pc,currentWord)
        if e !=nil {
            return e
        }
        e = currentWord.Do()
        // if it is a internal word increment the pc
        if currentWord.IsInternal() {
            fm.pc.inc()
        }
        if e != nil {
            if e == ErrExit {
                break
            }
            return e
        }
    }
        return e
	}
	fm.call = call
}

func (fm *ForthMachine) Exec(w Word) (err error) {
	fmt.Print("EXEC>")
	err = w.Do()
	return err
}

//Process the incoming
func (fm *ForthMachine) Process() (err error) {
	for {
		tok, empty := fm.NextToken()
		if empty {
			break
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
					//TODO add execute before every word
					fm.current.Add(fm.Get("call"))
					fm.current.Add(w)
				}
			} else {
				// spool into execution
				// TODO
				if w.IsLit() {
					err := fm.dStack.Push(w)
					if err != nil {
						return err
					}
				}
                fm.pc.Set(w,0)
				err  = w.Do()
                if err != nil {
                    return err
                }
			}
		}
	}
	return nil
}

