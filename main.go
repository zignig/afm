package main

import (
	"./fth"
)

func main() {
	options := fth.NewOptions("init", 32, 32, 4096)
	fm := fth.NewForthMachine(options)
	fm.Init()
	fm.Run()
}
