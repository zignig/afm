package fth

import (
	"fmt"
	"io"
)

// base forth machine structure

type Options struct {
	startword string
	rsize     int
	dsize     int
	mem       int
}

func NewOptions(startword string, rsize int, dsize int, mem int) (o Options) {
	o = Options{
		startword: startword,
		rsize:     rsize,
		dsize:     dsize,
		mem:       mem,
	}
	return o
}

type ForthMachine struct {
	Input  io.Writer
	Output io.Reader
	d      *ForthDictionary
	dStack Stack
	rStack Stack

	startword string
}

func NewForthMachine(o Options) (fm *ForthMachine) {
	fm = &ForthMachine{
		d:         NewForthDictionary(),
		rStack:    NewBaseStack(o.rsize),
		dStack:    NewBaseStack(o.dsize),
		startword: o.startword,
	}
	return fm
}

func (fm *ForthMachine) Add(w Word) (err error) {
	return fm.d.Add(w)
}

func (fm *ForthMachine) Words() {
	fm.d.Words()
}

func (fm *ForthMachine) Run() (e error) {
	fm.Words()
	fmt.Printf("Starting on %s\n", fm.startword)
	w, err := fm.d.Search(fm.startword)
	if err != nil {
		fmt.Println("boot error :", err)
		return err
	}
	w.Do()
	return
}
