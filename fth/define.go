package fth

import (
	"fmt"
)

// put the defining colon in own file
func (fm *ForthMachine) SetDef() {
	// return stack
	popRstack := NewBaseWord("<r")
	fm.Add(popRstack)

	def := NewBaseWord(":")
	def.Imm(true)
	fm.Add(def)

	// the magic define function
	defFunc := func() (e error) {
		fmt.Println("in define")
		name, empty := fm.NextToken() // grab the next colon
		if empty {
			return ErrNoMoreTokens
		}
		if debug {
			fmt.Println("test bind >", name)
		}
		newWord := NewBaseWord(name)
		fm.current = newWord
		fm.compile = true
		return
	}
	def.SetExec(defFunc)

	enddef := NewBaseWord(";")
	enddef.Imm(true)
	fm.Add(enddef)
	enddefFunc := func() (e error) {
		if debug {
			fmt.Println("end define")
		}
		if fm.current != nil {
			fm.Add(fm.current)
			fm.current.SetCode(fm.raw)
			fm.current.Add(popRstack)
		}
		fm.current = nil
		fm.compile = false
		return
	}
	enddef.SetExec(enddefFunc)
}
