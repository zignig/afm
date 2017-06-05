package afm

import (
	"bufio"
)

// base forth machine structure

type Options struct {
	Prompt    string
	startword string
	rsize     int
	dsize     int
	mem       int
	debug     bool
}

func NewOptions(prompt string, startword string, debug bool, rsize int, dsize int, mem int) (o Options) {
	o = Options{
		Prompt:    prompt,
		startword: startword,
		rsize:     rsize,
		dsize:     dsize,
		mem:       mem,
		debug:     debug,
	}
	return o
}

func DefaultOptions() (o Options) {
	o = Options{
		Prompt:    "<| ",
		startword: "init",
		rsize:     32,
		dsize:     32,
		mem:       4096,
		debug:     true,
	}
	return o
}

var debug bool

type ForthMachine struct {
	Input  chan string
	Output chan string
	d      *ForthDictionary
	dStack Stack
	rStack *Rstack

	prompt    string
	startword string
	Exit      bool

	// Options
	Options Options
	// completer
	Complete *Completer
	// state
	pc      *PCRef // program counter
	raw     string
	scanner *bufio.Scanner
	compile bool // true compiling / false intepreting
	// text processing
	current Word   // current word that the machine is currently working on
	token   string // current token
}

func NewForthMachine(o Options) (fm *ForthMachine) {
	fm = &ForthMachine{
		Input:     make(chan string, 1024),
		d:         NewForthDictionary(fm),
		dStack:    NewBaseStack("dstack", o.rsize),
		rStack:    NewRstack("rstack", o.dsize),
		startword: o.startword,
		prompt:    o.Prompt,
		Options:   o,
	}
	fm.Complete = NewCompleter(fm)
	return fm
}

// Expose some dictionary methods

func (fm *ForthMachine) Add(w Word) (err error) {
	return fm.d.Add(w)
}

func (fm *ForthMachine) Words() {
	fm.d.Words()
}
