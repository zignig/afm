package fth

// Dictionary
import (
	"errors"
	"fmt"
)

var (
	ErrWordNotFound = errors.New("word not found")
	ErrEndDict      = errors.New("end of dict")
)

type dictItem struct {
	name  string
	value Word
	//next  *dictItem
	prev *dictItem
}

type ForthDictionary struct {
	head *dictItem
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

func (fd *ForthDictionary) Words() {
	fmt.Println("WORDS")
	var current *dictItem
	current = fd.head
	for {
		if current == nil {
			break
		}
		fmt.Println(current.value.Name())
		current = current.prev
	}
	fmt.Println("END WORDS")
}

func (fd *ForthDictionary) Search(name string) (w Word, e error) {
	var current *dictItem
	current = fd.head
	for {
		if current == nil {
			return nil, ErrEndDict
		}
		if current.name == name {
			return current.value, nil
		}
		current = current.prev
	}
	return nil, ErrWordNotFound
}
