package main

import (
	"fmt"
)

// Tetromino building block of tetris
type Tetromino struct {
	shape [tSize][tSize]bool
	value byte
	x     uint8
	y     uint8
	field *Field
}

func (t *Tetromino) rotateLeft() *Tetromino {
	var nShape [tSize][tSize]bool

	for y, row := range t.shape {
		for x := range row {
			nShape[tSize-1-y][x] = t.shape[x][y]
		}
	}

	t.shape = nShape

	return t
}

func (t *Tetromino) rotateRight() *Tetromino {
	var nShape [tSize][tSize]bool

	for y, row := range t.shape {
		for x := range row {
			fmt.Println(x, y, y, 3-x, t.shape[x][y])

			nShape[tSize-1-x][y] = t.shape[x][y]
		}
	}

	t.shape = nShape

	return t
}

func (t *Tetromino) moveLeft() *Tetromino {
	t.remove()
	if t.canMove(-1, 0) {
		t.x--
	}
	t.place()

	return t
}

func (t *Tetromino) moveRight() *Tetromino {
	t.remove()
	if t.canMove(1, 0) {
		t.x++
	}
	t.place()

	return t
}

func (t *Tetromino) moveDown() *Tetromino {
	t.remove()
	t.y++
	t.place()
	return t
}

func (t *Tetromino) place() *Tetromino {
	x := int(t.x)
	y := int(t.y)

	for i, row := range t.shape {
		for j, value := range row {
			if value {
				t.field.state[i+y][j+x] = t.value
			}
		}
	}

	return t
}

func (t *Tetromino) remove() *Tetromino {
	x := int(t.x)
	y := int(t.y)

	for i, row := range t.shape {
		for j, value := range row {
			if value {
				t.field.state[i+y][j+x] = byte(0)
			}
		}
	}

	return t
}

func (t *Tetromino) canMove(x int, y int) bool {
	x = int(t.x) + x
	y = int(t.y) + y

	for i, row := range t.shape {
		for j, value := range row {
			if !value {
				continue
			}

			if i+y < 0 || i+y >= len(t.field.state) { // out of bounds top-bottom
				return false
			}

			if j+x < 0 || j+x >= len(t.field.state[i+y]) { // out of bounds left-right
				return false
			}

			if t.field.state[i+y][j+x] > 0 {
				return false
			}
		}
	}

	return true
}

func (t *Tetromino) allowedDown() bool {
	t.remove()
	b := t.canMove(0, 1)
	t.place()
	return b
}
