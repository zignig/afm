package afm

import ()

func (fm *ForthMachine) Extra() {
	// instrospacetion
	def := NewBaseWord("see")
	fm.Add(def)
	defFunc := func() (e error) {
		name, empty := fm.NextToken() // grab the next token
		if empty {
			return ErrNoMoreTokens
		}
		word, err := fm.d.Search(name)
		if err != nil {
			return err
		}
		fm.out(word, err)
		return
	}
	def.SetExec(defFunc)

	// pop off the stack

	pop := NewBaseWord(".")
	fm.Add(pop)
	popFunc := func() (e error) {
		w, e := fm.dStack.Pop()
		if e != nil {
			return e
		}
		fm.out("POP: ", w)
		return nil
	}
	pop.SetExec(popFunc)

	showStack := NewBaseWord(".s")
	fm.Add(showStack)
	showStackFunc := func() (e error) {
		fm.dStack.Show()
		return e
	}
	showStack.SetExec(showStackFunc)

	showRStack := NewBaseWord(".r")
	fm.Add(showRStack)
	showRStackFunc := func() (e error) {
		fm.rStack.Show()
		return e
	}
	showRStack.SetExec(showRStackFunc)

	// include externale file
	include := NewBaseWord("include")
	include.Imm(true)
	fm.Add(include)
	includeFunc := func() (e error) {
		name, empty := fm.NextToken() // grab the next colon
		if empty {
			return ErrNoMoreTokens
		}
		e = fm.LoadFile(name)
		if e != nil {
			return e
		}
		return
	}
	include.SetExec(includeFunc)
}
