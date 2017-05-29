package fth

// base forth machine structure

type ForthMachine struct {
	d      ForthDictionary
	dStack Stack
	rStack Stack
}

func NewForthMachine() (fm *ForthMachine) {
	fm = &ForthMachine{}
	return fm
}
