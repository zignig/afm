package afm

import ()

// put the defining colon in own file
func (fm *ForthMachine) SetDef() {
	// return stack
	popRstack := NewBaseWord("<r")
	fm.Add(popRstack)
	popRstackFunc := func() (e error) {
		fm.out("POP R STACK")
		pc, e := fm.rStack.Pop()
		if e != nil {
			return e
		}
		fm.pc.w = pc.Get(0)
		fm.pc.offset = pc.GetVal()
		return e
	}
	popRstack.SetExec(popRstackFunc)

	def := NewBaseWord(":")
	def.Imm(true)
	fm.Add(def)

	// the magic define function
	defFunc := func() (e error) {
		if debug {
			fm.out("in define")
		}
		name, empty := fm.NextToken() // grab the next token
		if empty {
			return ErrNoMoreTokens
		}
		if debug {
			fm.out("test bind >", name)
		}
		newWord := NewBaseWord(name)
		fm.current = newWord
		fm.compile = true
		return
	}
	def.SetExec(defFunc)

	// make word litteral
	lit := NewBaseWord("lit")
	fm.Add(lit)
	litFunc := func() (e error) {
		fm.current.Lit(true)
		return nil
	}
	lit.Imm(true)
	lit.SetExec(litFunc)

	// end compile
	enddef := NewBaseWord(";")
	enddef.Imm(true)
	fm.Add(enddef)
	enddefFunc := func() (e error) {
		if debug {
			fm.out("end define")
		}
		if fm.current != nil {
			if fm.current.IsImm() == false {
				fm.Add(fm.current)
			}
			fm.current.SetCode(fm.raw)
			fm.current.Add(popRstack)
			fm.current.SetExec(fm.call)
		}
		fm.compile = false
		return
	}
	enddef.SetExec(enddefFunc)
}
