package main

import (
	"./fth"
	"fmt"
)

func main() {
	fmt.Println("test")
	fth.Bip()
	o := fth.NewOptions(32, 32, 4096)
	fm := fth.NewForthMachine(o)
	fmt.Println(fm)
}
