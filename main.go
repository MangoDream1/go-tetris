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
	f := Field{}
	f.init()
	f.current.place()
	f.render()

	for {
		select {
		case <-ticker:
			gameover <- 0
			f.current.moveDown()
			f.tickActions()
		case event := <-keyEvents:
			gameover <- 0
			go handleInput(&f, event, quit)
		case <-quit:
			fmt.Println("QUITING")
			return
		case <-gameover:
			if !f.current.allowedDown() && f.isGameOver() {
				fmt.Println("Game over :(")
				return
			}
		}
	}
}
