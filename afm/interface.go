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
	fmt.Print(">| ")
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
