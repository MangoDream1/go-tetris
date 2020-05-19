package main

import (
	"math/rand"
	"time"
)

const fieldDropInZone = 1
const fieldHeight = 18 + fieldDropInZone
const fieldWidth = 12
const maxStored = 5
const winLinesCleared = 40

type duration struct {
	startTime time.Time
	endTime   time.Time
	duration  time.Duration
}

// Game holds the game state and logic components
type Game struct {
	field   [fieldHeight][fieldWidth]byte
	current Tetromino
	stored  Tetromino
	coming  []Tetromino
	cleared int8
	time    duration
}

func (f *Game) generateNew() *Game {
	i := rand.Intn(len(shapes))
	s := shapes[i]

	f.coming = append(f.coming, Tetromino{s.shape, s.value, fieldWidth/2 - tSize/4, 0, f})

	return f
}

func (f *Game) isWon() bool {
	hasWon := f.cleared >= winLinesCleared

	if hasWon {
		f.time.endTime = time.Now()
		f.time.duration = f.time.endTime.Sub(f.time.startTime).Round(time.Millisecond)

	}

	return hasWon
}

func (f *Game) isLost() bool {
	return !f.coming[0].canMove(0, 0, &f.coming[0].shape) // if next coming cant be rendered game over
}

func (f *Game) newCurrent() *Game {
	current, coming := f.coming[0], f.coming[1:]
	f.current = current
	f.coming = coming

	f.current.place()
	f.generateNew()

	return f
}

func (f *Game) storeCurrent() *Game {
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

func (f *Game) init() *Game {
	for i := 0; i < maxStored; i++ {
		f.generateNew()
	}

	f.newCurrent()
	f.field = [fieldHeight][fieldWidth]byte{}
	f.time.startTime = time.Now()

	return f
}

func (f *Game) tickActions() *Game {
	f.clearRowIfFilled(0)
	f.current.place()
	f.render()
	f.current.remove()
	return f
}

func (f *Game) clearRowIfFilled(iRow int) *Game {
	if iRow == fieldHeight {
		return f
	}

	row := f.field[iRow]
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
			row := f.field[i]
			f.field[i+1] = row
		}
		f.cleared++
	}

	return f.clearRowIfFilled(iRow + 1)
}
