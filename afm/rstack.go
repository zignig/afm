package afm

import (
	"fmt"
)

// Program counter reference
type PCRef struct {
	w      Word
	offset int
}

func (pc *PCRef) String() string {
	a := pc.w.Name() + " | " + string(pc.offset)
	return a
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

func (bs *Rstack) Show() {
	if bs.pos == 0 {
		fmt.Println("R stack empty")
		return
	}
	fmt.Println("R stack")
	for i := 0; i < bs.pos; i++ {
		fmt.Println(i, ":", bs.items[i])
	}
}
func (bs *Rstack) Pop() (pcr *PCRef, e error) {
	if bs.pos > 0 {
		bs.pos--
		pcr = bs.items[bs.pos]
		bs.items[bs.pos] = nil
		return pcr, nil
	}
	return nil, ErrStackEmpty
}

func (bs *Rstack) Push(pcr *PCRef) (e error) {
	if bs.pos == bs.size {
		return ErrStackFull
	}
	bs.items[bs.pos] = pcr
	bs.pos++
	return nil
}

func (bs *Rstack) Depth() (d int) {
	return bs.pos
}
