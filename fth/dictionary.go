package fth

// Dictionary

type dictItem struct {
	name string
	next *Word
	prev *Word
}

type ForthDictionary struct {
	head *Word
}
