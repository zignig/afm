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
	debug     bool
}

func NewOptions(startword string, debug bool, rsize int, dsize int, mem int) (o Options) {
	o = Options{
		startword: startword,
		rsize:     rsize,
		dsize:     dsize,
		mem:       mem,
		debug:     debug,
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
	debug     bool
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
	for {
		wl := GetLine()
		for _, j := range wl {
			fmt.Print(j, " : ")
			w, e := fm.d.Search(j)
			if fm.debug {
				fmt.Println(w, e)
			}
			if e == nil {
				w.Do()
			}
		}
	}

	fmt.Printf("Starting on %s\n", fm.startword)
	w, err := fm.d.Search(fm.startword)
	if err != nil {
		fmt.Println("boot error :", err)
		return err
	}
	w.Do()
	return
}
