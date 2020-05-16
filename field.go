package main

import (
	"math/rand"
)

const fieldDropInZone = 1
const fieldHeight = 18 + fieldDropInZone
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
	return !f.coming[0].canMove(0, 0, &f.coming[0].shape) // if next coming cant be rendered game over
}

func (f *Field) newCurrent() *Field {
	current, coming := f.coming[0], f.coming[1:]
	f.current = current
	f.coming = coming

	f.current.place()
	f.generateNew()

	return f
}

func (f *Field) storeCurrent() *Field {
	f.current.remove()
	if (f.stored == Tetromino{}) {
		f.newCurrent()
		f.stored = f.current
	} else {
		old := f.current
		f.current = f.stored
		f.stored = old
	}

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

func (f *Field) tickActions() *Field {
	f.clearRowIfFilled(0)
	f.current.place()
	f.render()
	f.current.remove()
	return f
}

func (f *Field) clearRowIfFilled(iRow int) *Field {
	if iRow == fieldHeight {
		return f
	}

	row := f.state[iRow]
	isFilled := true
	for _, value := range row {
		if value == byte(0) {
			isFilled = false
			break
		}
	}

	if isFilled {
		f.current.remove()
		for i := iRow - 1; i >= 0; i-- {
			row := f.state[i]
			f.state[i+1] = row
		}
	}

	return f.clearRowIfFilled(iRow + 1)
}
