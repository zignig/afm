package fth

// Stack, same kind of stack used for both

type Stack interface {
	Pop() (w Word, e error)
	Push(*Word) (e error)
	Depth() (d int)
}
