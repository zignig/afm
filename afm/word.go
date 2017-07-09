package afm

import (
	"errors"
	"fmt"
)

var (
	ErrOutOfBounds = errors.New("Out of bounds")
)

// Base interface
type execFunc func() (e error)

// TODO this interface is too big
type Word interface {
	Name() (s string)
	Do() (e error)
	Add(Word)
	SetExec(f execFunc) // for internal works
	SetExecExt (f execFunc) // for defined words
	SetCode(code string)
	Dump() (code string)
	IsImm() bool
	Imm(bool)
	IsLit() bool
	Lit(bool)
	Length() int
	Get(int) (w Word, err error)
	GetVal() (i int)
	SetVal(i int)
    IsInternal() bool
    Internal(bool)

}

type BaseWord struct {
	name      string
	immediate bool
	litteral  bool
    internal bool
	code      string
	words     []Word
	count     int
	exec      execFunc

	val int
	st  string
}

func (b *BaseWord) String() string {
	s := b.Name()
	if b.immediate {
		s += "|I"
	}
	if b.litteral {
		s += "|L"
	}
	s += fmt.Sprintf(" %d - ", b.count)
	if len(b.words) > 0 {
		for _, j := range b.words {
			s += fmt.Sprintf("%v ", j.Name())
		}
	}
	return s
}

func (b *BaseWord) GetVal() (i int) {
	return b.val
}

func (b *BaseWord) SetVal(i int) {
	b.val = i
}

func (b *BaseWord) Get(i int) (w Word, err error) {
	if i < len(b.words) {
		return b.words[i], nil
	}
	return nil, ErrOutOfBounds
}

func (b *BaseWord) Length() (i int) {
	return len(b.words)
}

func (b *BaseWord) Dump() (code string) {
	return b.code
}

func (b *BaseWord) SetCode(code string) {
	b.code = code
}

func (b *BaseWord) Internal(i bool) {
	b.internal= i
}

func (b *BaseWord) IsInternal() bool {
	return b.internal
}

func (b *BaseWord) Lit(i bool) {
	b.litteral = i
}

func (b *BaseWord) IsLit() bool {
	return b.litteral
}

func (b *BaseWord) Imm(i bool) {
	b.immediate = i
}

func (b *BaseWord) IsImm() bool {
	return b.immediate
}

func (b *BaseWord) SetExec(f execFunc) {
	b.exec = f
    b.internal = true
}

func (b *BaseWord) SetExecExt(f execFunc) {
	b.exec = f
    b.internal = false
}

func (b *BaseWord) Name() (s string) {
	return b.name
}

func (b *BaseWord) Add(w Word) {
	b.words = append(b.words, w)
}

func (b *BaseWord) Do() (e error) {
	// check if it has an internal GO function
	fmt.Println("DO ", b.Name(), " ")
	b.count++
	if b.exec != nil {
		e := b.exec()
		return e
	}
	return e
}

func NewBaseWord(name string) (w Word) {
	b := &BaseWord{
		name:  name,
		words: make([]Word, 0),
	}
	return b
}
