package fth

import (
	"bufio"
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
		Input:     make(chan string, 10),
		d:         NewForthDictionary(),
		rStack:    NewBaseStack("rstack", o.rsize),
		dStack:    NewBaseStack("dstack", o.dsize),
		startword: o.startword,
		prompt:    ">",
	}
	fm.Complete = NewCompleter(fm)
	// record the tail of the dictionary
	return fm
}

func (fm *ForthMachine) Add(w Word) (err error) {
	return fm.d.Add(w)
}

func (fm *ForthMachine) Words() {
	fm.d.Words()
}

func (fm *ForthMachine) Run(exit chan bool) (e error) {
	fm.Words()
	fmt.Printf("Starting on %s\n", fm.startword)
	w, err := fm.d.Search(fm.startword)
	if err != nil {
		fmt.Println("boot error :", err)
		return err
	}
	w.Do()
	for {
		if fm.exit {
			exit <- true
			break
		}
		select {
		case line := <-fm.Input:
			fmt.Println(line)
			fm.GetLine(line)
			fm.Process()
		}
		fmt.Println("ok")
	}
	return
}
