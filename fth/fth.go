package afm

import ()

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
			fm.Process()
		}
	}
	fm.out("EXIT MACHINE")
	return
}
