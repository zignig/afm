package afm

// Stack, this is the data stack
import (
	"errors"
	"fmt"
)

var (
	ErrStackFull  = errors.New("Stack Full")
	ErrStackEmpty = errors.New("Stack Empty")
)

type Stack interface {
	Pop() (w Word, e error)
	Push(Word) (e error)
	Depth() (d int)
	Show()
}

// Actual Stack

type BaseStack struct {
	name  string
	items []Word
	pos   int
	size  int
}

func NewBaseStack(name string, size int) (bs *BaseStack) {
	bs = &BaseStack{
		items: make([]Word, size),
		size:  size,
		name:  name,
		pos:   0,
	}
	return bs
}

func (bs *BaseStack) Show() {
	if bs.pos == 0 {
		fmt.Println(bs.name, " stack empty")
		return
	}
	fmt.Println(bs.name)
	for i := 0; i < bs.pos; i++ {
		fmt.Println(i, ":", bs.items[i])
	}
}

func (bs *BaseStack) Pop() (w Word, e error) {
	if bs.pos > 0 {
		bs.pos--
		w = bs.items[bs.pos]
		bs.items[bs.pos] = nil
		return w, nil
	}
	return nil, ErrStackEmpty
}

func (bs *BaseStack) Push(w Word) (e error) {
	if bs.pos == bs.size {
		return ErrStackFull
	}
	bs.items[bs.pos] = w
	bs.pos++
	return nil
}

func (bs *BaseStack) Depth() (d int) {
	fmt.Println("Depth")
	return len(bs.items)
}
