package fth

import (
	"fmt"
)

// put the defining colon in own file
func (fm *ForthMachine) Extra() {
	def := NewBaseWord("see")
	fm.Add(def)

	// the magic define function
	defFunc := func() (e error) {
		name, empty := fm.NextToken() // grab the next colon
		if empty {
			return ErrNoMoreTokens
		}
		word, err := fm.d.Search(name)
		fmt.Println(word, err)
		return
	}
	def.SetExec(defFunc)
}
