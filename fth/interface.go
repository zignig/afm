package fth

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrNoMoreTokens = errors.New("no more tokens")
)

func (fm *ForthMachine) GetLine(line string) {
	fm.raw = line
	fm.scanner = bufio.NewScanner(strings.NewReader(line))
	fm.scanner.Split(bufio.ScanWords)
}

func (fm *ForthMachine) NextToken() (s string, empty bool) {
	empty = !fm.scanner.Scan()
	s = fm.scanner.Text()
	return s, empty
}

func (fm *ForthMachine) LoadFile(name string) (err error) {
	fmt.Println("Load file ", name)
	_, err = os.Stat(name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fm.Input <- scanner.Text()
	}
	return
}

func (fm *ForthMachine) Process() (err error) {
	for {
		tok, empty := fm.NextToken()
		if empty {
			return ErrNoMoreTokens
		}
		w, err := fm.d.Search(tok)
		if debug {
			fmt.Println(w, err, tok)
		}
		if err != nil {
			fmt.Println(tok, "--", err)
			fm.compile = false
			return err
		} else {
			if fm.compile {
				if w.IsImm() {
					if debug {
						fmt.Println("imm function ", w.Name())
					}
					err = w.Do()
					if err != nil {
						return err
					}
				}
				if debug {
					fmt.Println("compile", fm.current, tok, w)
				}
				if fm.current != nil {
					fm.current.Add(w)
				}
			} else {
				err = w.Do()
				if err != nil {
					return err
				}
			}
		}
	}
}
