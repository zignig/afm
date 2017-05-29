package fth

// Stack, same kind of stack used for both
import (
	"fmt"
)

type Stack interface {
	Pop() (w *Word, e error)
	Push(*Word) (e error)
	Depth() (d int)
}

// Actual Stack

type BaseStack struct {
	items []*Word
}

func NewBaseStack(size int) (bs *BaseStack) {
	bs = &BaseStack{}
	bs.items = make([]*Word, size)
	return bs
}

func (bs *BaseStack) Pop() (w *Word, e error) {
	fmt.Println("Pop")
	return w, e
}

func (bs *BaseStack) Push(w *Word) (e error) {
	fmt.Println("Push")
	return e
}

func (bs *BaseStack) Depth() (d int) {
	fmt.Println("Depth")
	return len(bs.items)
}
