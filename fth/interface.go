package afm

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

func (fm *ForthMachine) out(s ...interface{}) {
	fmt.Print(">>>")
	fmt.Println(s...)
}

func (fm *ForthMachine) outf(s ...interface{}) {
	fmt.Println(s...)
}

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
	fm.out("Load file ", name)
	_, err = os.Stat(name)
	if err != nil {
		fm.out(err)
		return err
	}
	fm.out("Open File")
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return err
	}
	fm.out("Scan File")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fm.out(scanner.Text())
		fm.Input <- scanner.Text()
	}
	fm.out("Finished ", name)
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
			fm.out(w, err, tok)
		}
		if err != nil {
			fm.out(tok, "--", err)
			fm.compile = false
			return err
		} else {
			if fm.compile {
				if w.IsImm() {
					if debug {
						fm.out("imm function ", w.Name())
					}
					err = w.Do()
					if err != nil {
						return err
					}
				}
				if debug {
					fm.out("compile", fm.current, tok, w)
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
