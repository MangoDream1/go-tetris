package main

// Tetromino building block of tetris
type Tetromino struct {
	shape [tSize][tSize]bool
	value byte
	x     int8
	y     int8
	game  *Game
}

func (t *Tetromino) rotateLeft() *Tetromino {
	var proposedShape [tSize][tSize]bool
	for iRow, row := range t.shape {
		for i := range row {
			proposedShape[tSize-1-iRow][i] = t.shape[i][iRow]
		}
	}

	fixShapeLeft(&proposedShape)
	fixShapeUp(&proposedShape)

	if !t.canMove(0, 0, &proposedShape) {
		return t
	}

	t.shape = proposedShape
	return t
}

func (t *Tetromino) rotateRight() *Tetromino {
	var proposedShape [tSize][tSize]bool
	for iRow, row := range t.shape {
		for i := range row {
			proposedShape[iRow][tSize-1-i] = t.shape[i][iRow]
		}
	}

	fixShapeLeft(&proposedShape)
	fixShapeUp(&proposedShape)

	if !t.canMove(0, 0, &proposedShape) {
		return t
	}

	t.shape = proposedShape
	return t
}

func (t *Tetromino) moveLeft() *Tetromino {
	if t.canMove(-1, 0, &t.shape) {
		t.x--
	}

	return t
}

func (t *Tetromino) moveRight() *Tetromino {
	if t.canMove(1, 0, &t.shape) {
		t.x++
	}

	return t
}

func (t *Tetromino) moveDown() *Tetromino {
	if !t.allowedDown() {
		t.place()
		t.game.newCurrent()
	} else {
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
				t.game.field[i+y][j+x] = t.value
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
				t.game.field[i+y][j+x] = byte(0)
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

			if i+y < 0 || i+y >= len(t.game.field) { // out of bounds top-bottom
				return false
			}

			if j+x < 0 || j+x >= len(t.game.field[i+y]) { // out of bounds left-right
				return false
			}

			if t.game.field[i+y][j+x] > 0 {
				return false
			}
		}
	}

	return true
}

func (t *Tetromino) allowedDown() bool {
	b := t.canMove(0, 1, &t.shape)
	return b
}
