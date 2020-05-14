package main

// Tetromino building block of tetris
type Tetromino struct {
	shape [tSize][tSize]bool
	value byte
	x     uint8
	y     uint8
	field *Field
}

// FIXME: check if allowed to turn
// TODO: remove bottom row if filled

func (t *Tetromino) rotateLeft() *Tetromino {
	t.remove()
	var nShape [tSize][tSize]bool

	for iRow, row := range t.shape {
		for i := range row {
			nShape[iRow][tSize-1-i] = t.shape[i][iRow]
		}
	}

	t.shape = nShape
	t.place()
	return t
}

func (t *Tetromino) rotateRight() *Tetromino {
	t.remove()
	var nShape [tSize][tSize]bool

	for iRow, row := range t.shape {
		for i := range row {
			nShape[tSize-1-iRow][i] = t.shape[i][iRow]
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

	return t
}

func (t *Tetromino) moveRight() *Tetromino {
	t.remove()
	if t.canMove(1, 0) {
		t.x++
	}

	return t
}

func (t *Tetromino) moveDown() *Tetromino {
	if !t.allowedDown() {
		t.place()
		t.field.newCurrent()
	} else {
		t.remove()
		t.y++
	}

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
	return b
}
