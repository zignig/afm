package fth

// Dictionary
import (
	"errors"
	"fmt"
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
	head *dictItem
	stem *dictItem
}

func NewForthDictionary() (fd *ForthDictionary) {
	fd = &ForthDictionary{}
	return fd
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
	var current *dictItem
	current = fd.head
	for {
		if current == nil {
			break
		}
		val := current.value.Dump()
		if len(val) > 0 {
			fmt.Print(val)
		}
		current = current.prev
	}
	fmt.Println()
}

func (fd *ForthDictionary) Find(line []rune, pos int) (l []string) {
	var current *dictItem
	current = fd.head
	str := string(line)
	list := make([]string, 0)
	for {
		if current == nil {
			break
		}
		if strings.HasPrefix(current.name, str) {
			list = append(list, current.name)
		}
		current = current.prev
	}
	return list
}

func (fd *ForthDictionary) Words() {
	var current *dictItem
	current = fd.head
	for {
		if current == nil {
			break
		}
		if debug {
			fmt.Println(current.value)
		} else {
			fmt.Print(current.value.Name(), " ")
		}
		current = current.prev
	}
	fmt.Println()
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
