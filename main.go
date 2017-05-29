package main

import (
	"./fth"
)

func main() {
	options := fth.NewOptions("init", true, 32, 32, 4096)
	fm := fth.NewForthMachine(options)
	fm.Init()
	fm.Run()
}
