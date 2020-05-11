package main

import (
	"math/rand"
)

const fieldHeight = 18
const fieldWidth = 12
const maxStored = 5

// const buffer = tSize

type Field struct {
	state   [fieldHeight][fieldWidth]byte
	current Tetromino
	stored  Tetromino
	coming  []Tetromino
}

func (f *Field) removeCurrent() *Field {
	x := int(f.current.x)
	y := int(f.current.y)

	for i, row := range f.current.shape {
		for j, value := range row {
			if value {
				f.state[i+y][j+x] = byte(0)
			}
		}
	}

	return f
}

func (f *Field) placeCurrent() *Field {
	x := int(f.current.x)
	y := int(f.current.y)

	for i, row := range f.current.shape {
		for j, value := range row {
			if value {
				f.state[i+y][j+x] = f.current.value
			}
		}
	}

	return f
}

func (f *Field) generateNew() *Field {
	i := rand.Intn(len(shapes))
	s := shapes[i]

	f.coming = append(f.coming, Tetromino{s.shape, s.value, fieldWidth/2 - tSize/2, 0})

	return f
}

func (f *Field) setCurrent() *Field {
	current, coming := f.coming[0], f.coming[1:]
	f.current = current
	f.coming = coming

	f.generateNew()

	return f
}

func (f *Field) storeCurrent() *Field {
	if (f.stored == Tetromino{}) {
		f.setCurrent()
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

	f.setCurrent()
	f.state = [fieldHeight][fieldWidth]byte{}

	return f
}
