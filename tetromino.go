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
}

func indexing(x int8, y int8) int8 {
	return x + y*tSize
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
	t.x--
	return t
}

func (t *Tetromino) moveRight() *Tetromino {
	t.x++
	return t
}

func (t *Tetromino) moveDown() *Tetromino {
	t.y++
	return t
}
