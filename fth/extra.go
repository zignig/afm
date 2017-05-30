package fth

import (
	"fmt"
)

// put the defining colon in own file
func (fm *ForthMachine) Extra() {
	// instrospacetion
	def := NewBaseWord("see")
	fm.Add(def)
	defFunc := func() (e error) {
		name, empty := fm.NextToken() // grab the next colon
		if empty {
			return ErrNoMoreTokens
		}
		word, err := fm.d.Search(name)
		if err != nil {
			return err
		}
		fmt.Println(word, err)
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
		fmt.Println(w)
		return nil
	}
	pop.SetExec(popFunc)

	// include externale file
	include := NewBaseWord("include")
	fm.Add(include)
}
