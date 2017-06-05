package afm

// Dictionary
import (
	"errors"
	"strings"
)

var (
	ErrWordNotFound = errors.New("word not found")
	ErrEndDict      = errors.New("end of dict")
)

type dictItem struct {
	name  string
	value Word
	next  *dictItem
	prev  *dictItem
}

type ForthDictionary struct {
	fm   *ForthMachine
	head *dictItem
	stem *dictItem
}

func NewForthDictionary(fm *ForthMachine) (fd *ForthDictionary) {
	fd = &ForthDictionary{fm: fm}
	return fd
}

func (fd *ForthDictionary) out(s ...interface{}) {
	fd.fm.out(s...)
}

func (fd *ForthDictionary) Add(w Word) (e error) {
	// should search to see if the word is a redef
	newDict := &dictItem{
		name:  w.Name(),
		value: w,
		prev:  fd.head,
	}
	fd.head = newDict
	return
}

func (fd *ForthDictionary) dump() {
	fd.out()
	var current *dictItem
	current = fd.head
	for {
		if current == nil {
			break
		}
		val := current.value.Dump()
		if len(val) > 0 {
			fd.fm.out(val)
		}
		current = current.prev
	}
	fd.out()
}

// auto complete
func (fd *ForthDictionary) Find(line []rune, pos int) (newLine [][]rune, length int) {
	var current *dictItem
	current = fd.head
	str := string(line)
	list := make([][]rune, 0)
	for {
		if current == nil {
			break
		}
		if strings.HasPrefix(current.name, str) {
			list = append(list, []rune(current.name[pos:]))
		}
		current = current.prev
	}
	return list, 0
}

func (fd *ForthDictionary) Words() {
	var current *dictItem
	current = fd.head
	for {
		if current == nil {
			break
		}
		if debug {
			fd.out(current.value)
		} else {
			fd.out(current.value.Name(), " ")
		}
		current = current.prev
	}
	fd.out()
}

func (fd *ForthDictionary) Search(name string) (w Word, e error) {
	var current *dictItem
	current = fd.head
	for {
		if current == nil {
			return nil, ErrWordNotFound
		}
		if current.name == name {
			return current.value, nil
		}
		current = current.prev
	}
	return nil, ErrWordNotFound
}
