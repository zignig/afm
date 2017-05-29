package fth

import (
	"fmt"
)

// Base interface
type execFunc func() (e error)

type Word interface {
	Name() (s string)
	Do() (e error)
	SetExec(f execFunc)
}

type BaseWord struct {
	name      string
	immediate bool
	code      string
	words     []Word
	count     int
	exec      execFunc
}

func (b *BaseWord) SetExec(f execFunc) {
	b.exec = f
}

func (b *BaseWord) Name() (s string) {
	return b.name
}

func (b *BaseWord) Do() (e error) {
	// check if it has an internal GO function
	b.count++
	if b.exec != nil {
		b.exec()
		return
	}
	fmt.Println("interpret")
	return e
}

func NewBaseWord(name string) (w Word) {
	b := &BaseWord{
		name: name,
	}
	return b
}
