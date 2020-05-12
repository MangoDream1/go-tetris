package main

import (
	"math/rand"
)

const fieldHeight = 18
const fieldWidth = 12
const maxStored = 5

type Field struct {
	state   [fieldHeight][fieldWidth]byte
	current Tetromino
	stored  Tetromino
	coming  []Tetromino
}

func (f *Field) generateNew() *Field {
	i := rand.Intn(len(shapes))
	s := shapes[i]

	f.coming = append(f.coming, Tetromino{s.shape, s.value, fieldWidth/2 - tSize/2, 0, f})

	return f
}

func (f *Field) isGameOver() bool {
	return !f.coming[0].canMove(0, 0) // if next coming cant be rendered game over
}

func (f *Field) newCurrent() *Field {
	current, coming := f.coming[0], f.coming[1:]
	f.current = current
	f.coming = coming

	f.generateNew()

	return f
}

func (f *Field) storeCurrent() *Field {
	f.current.remove()
	if (f.stored == Tetromino{}) {
		f.newCurrent()
	} else {
		f.current = f.stored
	}

	f.stored = f.current
	return f
}

func (f *Field) init() *Field {
	for i := 0; i < maxStored; i++ {
		f.generateNew()
	}

	f.newCurrent()
	f.state = [fieldHeight][fieldWidth]byte{}

	return f
}
