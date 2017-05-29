package fth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLine() (words []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">")
	text, _ := reader.ReadString('\n')
	val := strings.Split(strings.TrimSpace(text), " ")
	for i, j := range val {
		fmt.Println(i, ">", j, "<")
	}
	return val
}
