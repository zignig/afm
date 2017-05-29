package main

import (
	"./fth"
	"fmt"
)

func main() {
	fmt.Println("test")
	fth.Bip()
	fm := fth.NewForthMachine()
	fmt.Println(fm)
}
