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

func (fm *ForthMachine) GetLine() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(fm.prompt)
	text, _ := reader.ReadString('\n')
	fm.raw = text
	fm.scanner = bufio.NewScanner(strings.NewReader(text))
	fm.scanner.Split(bufio.ScanWords)
}

func (fm *ForthMachine) NextToken() (s string, empty bool) {
	empty = !fm.scanner.Scan()
	s = fm.scanner.Text()
	return s, empty
}

func (fm *ForthMachine) Process() (err error) {
	for {
		tok, empty := fm.NextToken()
		if empty {
			return ErrNoMoreTokens
		}
		if debug {
			fmt.Println(tok)
		}
		w, err := fm.d.Search(tok)
		if err != nil {
			fmt.Println(tok, "--", err)
			fm.compile = false
			return err
		} else {
			if fm.compile {
				if w.IsImm() {
					fmt.Println("imm function ", w.Name())
					err = w.Do()
					if err != nil {
						return err
					}
				}
				fmt.Println("compile", fm.current, tok, w)
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
