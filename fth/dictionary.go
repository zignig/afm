package fth

// Dictionary

type dictItem struct {
	name  string
	value Word
	next  *dictItem
	prev  *dictItem
}

type ForthDictionary struct {
	head *dictItem
}

func NewForthDictionary() (fd *ForthDictionary) {
	fd = &ForthDictionary{}
	return fd
}

func (fd *ForthDictionary) Add(w Word) (e error) {
	newDict := &dictItem{
		name:  w.Name(),
		value: w,
		prev:  fd.head,
	}
	fd.head = newDict
	return
}

func (fd *ForthDictionary) Search(name string) (w Word, e error) {
	return
}
