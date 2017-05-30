package fth

import (
	"fmt"
)

// Base interface
type execFunc func() (e error)

type Word interface {
	Name() (s string)
	Do() (e error)
	Add(Word)
	SetExec(f execFunc)
	IsImm() bool
	Imm(bool)
}

type BaseWord struct {
	name      string
	immediate bool
	code      string
	words     []Word
	count     int
	exec      execFunc
}

func (b *BaseWord) Imm(i bool) {
	b.immediate = i
}

func (b *BaseWord) IsImm() bool {
	return b.immediate
}

func (b *BaseWord) SetExec(f execFunc) {
	b.exec = f
}

func (b *BaseWord) Name() (s string) {
	return b.name
}

func (b *BaseWord) Add(w Word) {
	b.words = append(b.words, w)
}

func (b *BaseWord) Do() (e error) {
	// check if it has an internal GO function
	b.count++
	if b.exec != nil {
		b.exec()
		return
	}
	fmt.Println("run")
	fmt.Println(b)
	return e
}

func NewBaseWord(name string) (w Word) {
	b := &BaseWord{
		name:  name,
		words: make([]Word, 0),
	}
	return b
}
