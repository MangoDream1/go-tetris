package main

import (
	"fmt"
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

func tempPlaceCurrent(f *Field) [fieldHeight][fieldWidth]byte {
	return f.state
}

func (f *Field) placeCurrent() *Field {
	f.state = tempPlaceCurrent(f)
	return f
}

func (f *Field) generateNew() *Field {
	i := rand.Intn(len(shapes))
	s := shapes[i]

	f.coming = append(f.coming, Tetromino{s.shape, s.value, 0, 0})

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
	for i := 0; i < maxStored-1; i++ {
		f.generateNew()
	}

	f.setCurrent()
	f.state = [fieldHeight][fieldWidth]byte{}

	return f
}

func (f *Field) render() *Field {
	output := make([][]byte, fieldHeight+1)

	for i, row := range f.state {
		output[i] = make([]byte, fieldWidth+2)

		// define edges
		output[i][0] = '|'
		output[i][fieldWidth+1] = '|'

		for j, b := range row {
			if b == byte(0) {
				output[i][j+1] = ' '
				continue
			}

			output[i][j+1] = b
		}
	}

	// add bottom
	output[fieldHeight] = make([]byte, fieldWidth+2)
	for i := 0; i < fieldWidth+2; i++ {
		output[fieldHeight][i] = '='
	}

	fmt.Printf("\033[2;0H") // move cursor to second row and draw from there
	for _, row := range output {
		fmt.Println(string(row))
	}

	return f
}
