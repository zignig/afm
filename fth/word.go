package fth

import (
	"fmt"
)

// Base interface

type Word interface {
	Do() (e error)
}

type BaseWord struct {
	name  string
	code  string
	words []Word
	count int
	exec  *func()
}

func (b *BaseWord) Do() (e error) {
	fmt.Println("NOTHING HERE ", b.name)
	return e
}
