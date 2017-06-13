package afm

import (
	"errors"
)

var (
	ErrWordOverflow = errors.New("Too many words")
)

// Program counter reference
type PCRef struct {
	w      Word
	offset int
}

func (pc *PCRef) Set(w Word, offset int) (err error) {
	w, err = w.Get(0)
	if err != nil {
		return err
	}
	pc.w = w
	pc.offset = offset
	return
}

func (pc *PCRef) Get() (w Word, err error) {
	w, err = pc.w.Get(pc.offset)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (pc *PCRef) String() string {
	a := pc.w.Name() + " | " + string(pc.offset)
	return a
}

func (pc *PCRef) inc() (e error) {
	if pc.offset < pc.w.Length() {
		pc.offset++
		return
	}
	return ErrWordOverflow
}

// wrap a pc in a word so it can go into the rstack
func (pc *PCRef) wrap() (w Word) {
	w = NewBaseWord(pc.w.Name())
	// set the target word as the first ref
	w.Add(pc.w)
	// set the val as the offset
	w.SetVal(pc.offset)
	return
}
