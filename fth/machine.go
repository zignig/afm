package fth

import (
	"bufio"
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

var debug bool

type ForthMachine struct {
	Input  chan string
	Output io.Reader
	d      *ForthDictionary
	dStack Stack
	rStack Stack

	prompt    string
	startword string
	exit      bool

	// completer
	Complete *Completer
	// state
	raw     string
	scanner *bufio.Scanner
	compile bool // true compiling / false intepreting
	// text processing
	current Word // current word that the machine is currently working on
}

func NewForthMachine(o Options) (fm *ForthMachine) {
	fm = &ForthMachine{
		Input:     make(chan string, 1000),
		d:         NewForthDictionary(),
		rStack:    NewBaseStack("rstack", o.rsize),
		dStack:    NewBaseStack("dstack", o.dsize),
		startword: o.startword,
		prompt:    ">",
	}
	fm.Complete = NewCompleter(fm)
	return fm
}

func (fm *ForthMachine) Add(w Word) (err error) {
	return fm.d.Add(w)
}

func (fm *ForthMachine) Words() {
	fm.d.Words()
}
