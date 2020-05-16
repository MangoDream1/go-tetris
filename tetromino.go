package main

// Tetromino building block of tetris
type Tetromino struct {
	shape [tSize][tSize]bool
	value byte
	x     int8
	y     int8
	field *Field
}

func (t *Tetromino) rotateLeft() *Tetromino {
	t.remove()

	var proposedShape [tSize][tSize]bool
	for iRow, row := range t.shape {
		for i := range row {
			proposedShape[tSize-1-iRow][i] = t.shape[i][iRow]
		}
	}

	if !t.canMove(0, 0, &proposedShape) {
		return t
	}

	t.shape = proposedShape
	return t
}

func (t *Tetromino) rotateRight() *Tetromino {
	t.remove()

	var proposedShape [tSize][tSize]bool
	for iRow, row := range t.shape {
		for i := range row {
			proposedShape[iRow][tSize-1-i] = t.shape[i][iRow]
		}
	}

	if !t.canMove(0, 0, &proposedShape) {
		return t
	}

	t.shape = proposedShape
	return t
}

func (t *Tetromino) moveLeft() *Tetromino {
	t.remove()
	if t.canMove(-1, 0, &t.shape) {
		t.x--
	}

	return t
}

func (t *Tetromino) moveRight() *Tetromino {
	t.remove()
	if t.canMove(1, 0, &t.shape) {
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

func (t *Tetromino) hardDrop() *Tetromino {
	for t.allowedDown() {
		t.moveDown()
	}
	t.moveDown() // set in place
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

func (t *Tetromino) canMove(x int, y int, proposedShape *[tSize][tSize]bool) bool {
	x = int(t.x) + x
	y = int(t.y) + y

	for i, row := range proposedShape {
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
	b := t.canMove(0, 1, &t.shape)
	return b
}
