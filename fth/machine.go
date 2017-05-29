package fth

import (
	"io"
)

// base forth machine structure

type Options struct {
	rsize int
	dsize int
	mem   int
}

func NewOptions(rsize int, dsize int, mem int) (o Options) {
	o = Options{
		rsize: rsize,
		dsize: dsize,
		mem:   mem,
	}
	return o
}

type ForthMachine struct {
	Input  io.Writer
	Output io.Reader
	d      ForthDictionary
	dStack Stack
	rStack Stack
}

func NewForthMachine(o Options) (fm *ForthMachine) {
	fm = &ForthMachine{}
	fm.rStack = NewBaseStack(o.rsize)
	fm.dStack = NewBaseStack(o.dsize)
	return fm
}
