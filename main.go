package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

const tSize = 4

func main() {
	quit := make(chan int)
	gameover := make(chan int, 1)

	keyEvents := setupKeyboard()
	defer func() {
		_ = keyboard.Close()
	}()

	rand.Seed(time.Now().UnixNano())
	ticker := time.Tick(time.Second)

	// init game
	f := Game{}
	f.init()
	f.tickActions()

	for {
		select {
		case <-ticker:
			f.current.moveDown()
			gameover <- 0
			f.tickActions()
		case event := <-keyEvents:
			gameover <- 0
			go handleInput(&f, event, quit)
		case <-quit:
			fmt.Println("QUITING")
			return
		case <-gameover:
			f.current.remove()

			if f.isWon() {
				fmt.Printf("Game won! Congrats! %v", f.time.duration.String())
				return
			}

			if !f.current.allowedDown() && f.isLost() {
				fmt.Println("Game over :(")
				return
			}
		}
	}
}
