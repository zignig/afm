package fth

import (
	"fmt"
)

// Base interface

type Word interface {
	Name() (s string)
	Do() (e error)
}

type BaseWord struct {
	name  string
	code  string
	words []Word
	count int
	exec  func()
}

func (b *BaseWord) Name() (s string) {
	return b.name
}

func (b *BaseWord) Do() (e error) {
	fmt.Println("NOTHING HERE ", b.name)
	if b.exec != nil {
		b.exec()
	}
	return e
}
