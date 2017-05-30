package fth

import (
	"bytes"
	"fmt"
)

type PrefixCompleterInterface interface {
	Print(prefix string, level int, buf *bytes.Buffer)
	Do(line []rune, pos int) (newLine [][]rune, length int)
	GetName() []rune
	GetChildren() []PrefixCompleterInterface
	SetChildren(children []PrefixCompleterInterface)
}

type Completer struct {
	fm *ForthMachine
}

func NewCompleter(fm *ForthMachine) (c *Completer) {
	c = &Completer{
		fm: fm,
	}
	return
}

func (c *Completer) Print(prefix string, level int, buf *bytes.Buffer) {
}

func (c *Completer) Do(line []rune, pos int) (newLine [][]rune, length int) {
	fmt.Println(c.fm.d.List())
	fmt.Println(line, pos)
	return
}

func (c *Completer) GetName() []rune {
	return make([]rune, 0)
}

func (c *Completer) GetChildren() []PrefixCompleterInterface {
	return make([]PrefixCompleterInterface, 0)
}

func (c *Completer) SetChildren(children []PrefixCompleterInterface) {
	return
}
