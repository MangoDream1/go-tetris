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
			if !f.current.allowedDown() && f.isGameOver() {
				fmt.Println("Game over :(")
				return
			}

			f.current.moveDown()
			f.current.place()
			f.render()
		case event := <-keyEvents:
			if !f.current.allowedDown() && f.isGameOver() {
				fmt.Println("Game over :(")
				return
			}

			go handleInput(&f, event, quit)
		case <-quit:
			fmt.Println("QUITING")
			return
		}
	}
}
