package fth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (fm *ForthMachine) GetLine() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(fm.prompt)
	text, _ := reader.ReadString('\n')
	val := strings.Fields(text)
	fm.raw = text
	fm.tokens = val
	fm.tokenp = 0
}

func (fm *ForthMachine) Process() {
	for i, j := range fm.tokens {
		if fm.debug {
			fmt.Println(i, j)
		}
		w, err := fm.d.Search(j)
		if err != nil {
			fmt.Println(j, "--", err)
		} else {
			if fm.compile {
				if w.IsImm() {
					fmt.Println("im function")
					w.Do()
				}
				fmt.Println("do stuff with ", j)
			} else {
				w.Do()
			}
		}
	}
}
