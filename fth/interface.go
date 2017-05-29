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
	fm.raw = text
}

func (fm *ForthMachine) Process() {
	scanner := bufio.NewScanner(strings.NewReader(fm.raw))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		j := scanner.Text()
		if fm.debug {
			fmt.Println(j)
		}
		w, err := fm.d.Search(j)
		if err != nil {
			fmt.Println(j, "--", err)
			fm.compile = false
		} else {
			if fm.compile {
				if w.IsImm() {
					fmt.Println("im function")
					w.Do()
				}
				fmt.Println("compile", j)
			} else {
				w.Do()
			}
		}
	}
}
