package fth

import ()

type ref struct {
	w      *Word
	offset int
}

type Rstack struct {
	name  string
	items []*ref
	size  int
}

func NewRstack(name string, depth int) (r *Rstack) {
	r = &Rstack{
		name:  name,
		items: make([]*ref, depth),
		size:  depth,
	}
	return
}

func (bs *Rstack) Pop() (w *Word, e error) {
	return nil, ErrStackEmpty
}

func (bs *Rstack) Push(w *Word) (e error) {
	return nil
}

func (bs *Rstack) Depth() (d int) {
	return len(bs.items)
}
