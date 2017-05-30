package fth

import (
	"fmt"
)

// put the defining colon in own file
func (fm *ForthMachine) SetDef() {
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
		fmt.Println("test bind >", name)
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
		fmt.Println("end define")
		fm.Add(fm.current)
		fm.compile = false
		return
	}
	enddef.SetExec(enddefFunc)
}
