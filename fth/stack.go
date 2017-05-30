package fth

// Stack, same kind of stack used for both
import (
	"errors"
	"fmt"
)

var (
	ErrStackFull  = errors.New("Stack Full")
	ErrStackEmpty = errors.New("Stack Empty")
)

type Stack interface {
	Pop() (w *Word, e error)
	Push(*Word) (e error)
	Depth() (d int)
}

// Actual Stack

type BaseStack struct {
	name  string
	items []*Word
	size  int
}

func NewBaseStack(name string, size int) (bs *BaseStack) {
	bs = &BaseStack{
		items: make([]*Word, size),
		size:  size,
		name:  name,
	}
	return bs
}

func (bs *BaseStack) Pop() (w *Word, e error) {
	fmt.Println("Pop")
	if len(bs.items) > 0 {
		w = bs.items[len(bs.items)-1]
		bs.items = bs.items[:len(bs.items)-1]
		return w, nil
	}
	return nil, ErrStackEmpty
}

func (bs *BaseStack) Push(w *Word) (e error) {
	fmt.Println("Push")
	if len(bs.items) == bs.size {
		return ErrStackFull
	}
	bs.items = append(bs.items, w)
	return nil
}

func (bs *BaseStack) Depth() (d int) {
	fmt.Println("Depth")
	return len(bs.items)
}
