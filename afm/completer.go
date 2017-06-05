package afm

import (
//	"fmt"
//	"strings"
)

type AutoCompleter interface {
	Do(line []rune, pos int) (newLine [][]rune, length int)
}

type Completer struct {
	fm *ForthMachine
}

func NewCompleter(fm *ForthMachine) (c *Completer) {
	c = &Completer{
		fm: fm,
	}
	return
}

func (c *Completer) Do(line []rune, pos int) (newLine [][]rune, length int) {
	//typedSoFar := string(line[:pos])
	//spacePos := strings.IndexByte(typedSoFar, ' ')
	//entrim := line[spacePos+1:]
	//if spacePos > 0 {
	//fmt.Printf(">>%v<<\n", string(entrim))
	//}
	newLine, length = c.fm.d.Find(line, pos)
	return
}
