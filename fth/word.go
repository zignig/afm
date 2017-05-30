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
	SetCode(code string)
	Dump() (code string)
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

func (b *BaseWord) String() string {
	s := b.Name()
	if b.immediate {
		s += "(IMM)"
	}
	s += "\ncode > " + b.code
	s += "\n"
	if len(b.words) > 0 {
		for _, j := range b.words {
			s += fmt.Sprintf("%v ", j.Name())
		}
	}
	return s
}

func (b *BaseWord) Dump() (code string) {
	return b.code
}

func (b *BaseWord) SetCode(code string) {
	b.code = code
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
	fmt.Println("runing >", b.name)
	return e
}

func NewBaseWord(name string) (w Word) {
	b := &BaseWord{
		name:  name,
		words: make([]Word, 0),
	}
	return b
}
