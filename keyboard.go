package main

import (
	"github.com/eiannone/keyboard"
)

func setupKeyboard() <-chan keyboard.KeyEvent {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}

	return keysEvents
}

func handleInput(f *Field, event keyboard.KeyEvent, quit chan int) {
	if event.Err != nil {
		panic(event.Err)
	}

	if event.Key == keyboard.KeyEsc || event.Key == keyboard.KeyCtrlC {
		quit <- 0
		return
	}

	if event.Rune == 'c' {
		f.storeCurrent()
	}

	if event.Rune == 'z' {
		f.current.rotateLeft()
	}

	if event.Key == keyboard.KeyArrowUp {
		f.current.rotateRight()
	}

	if event.Key == keyboard.KeyArrowLeft {
		f.current.moveLeft()
	}

	if event.Key == keyboard.KeyArrowRight {
		f.current.moveRight()
	}

	if event.Key == keyboard.KeyArrowDown {
		f.current.moveDown()
	}

	if event.Key == keyboard.KeySpace {
		f.current.moveDown() // hard drop
	}

	f.tickActions()
}
