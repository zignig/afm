package fth

import (
	"fmt"
)

// put the defining colon in own file
func (fm *ForthMachine) SetDef() {
	fmt.Println("define colon")
	def := NewBaseWord(":")
	def.Imm(true)
	fm.Add(def)

	// the magic define function
	defFunc := func() (e error) {
		fmt.Println("in define")
		fm.NextToken()         // chomp the colon
		name := fm.NextToken() // grab the next colon
		fmt.Println("test bind >", name)
		newWord := NewBaseWord(name)
		fm.current = newWord
		fm.compile = true
		fm.Add(newWord)
		return
	}
	def.SetExec(defFunc)

	enddef := NewBaseWord(";")
	enddef.Imm(true)
	fm.Add(enddef)
	enddefFunc := func() (e error) {
		fmt.Println("end define")
		fm.compile = false
		return
	}
	enddef.SetExec(enddefFunc)
}
