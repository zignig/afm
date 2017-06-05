package afm

import ()

type PCRef struct {
	w      *Word
	offset int
}

type Rstack struct {
	name  string
	items []*PCRef
	pos   int
	size  int
}

func NewRstack(name string, depth int) (r *Rstack) {
	r = &Rstack{
		name:  name,
		items: make([]*PCRef, depth),
		size:  depth,
		pos:   0,
	}
	return
}

func (bs *Rstack) Pop() (pcr *PCRef, e error) {
	return nil, ErrStackEmpty
}

func (bs *Rstack) Push(pcr *PCRef) (e error) {
	return nil
}

func (bs *Rstack) Depth() (d int) {
	return len(bs.items)
}
