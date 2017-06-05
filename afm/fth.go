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
	w.Do()
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
		fmt.Println(err)
	}
	fm.out("EXIT MACHINE")
	return
}

func (fm *ForthMachine) Process() (err error) {
	for {
		tok, empty := fm.NextToken()
		if empty {
			return ErrNoMoreTokens
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
				if w.IsLit() {
					fmt.Println("CHUCK ME ON THE STACK")
					err = fm.dStack.Push(w)
					if err != nil {
						return err
					}
					continue
				}
				err = w.Do()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
