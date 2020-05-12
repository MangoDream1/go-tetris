package main

import (
	"fmt"
	"math/rand"
	"time"
)

const tSize = 4

func main() {
	rand.Seed(time.Now().UnixNano())

	f := Field{}
	f.init()
	f.current.place()
	f.render()

	ticker := time.Tick(time.Second / 5)

	f.storeCurrent()

	for true {
		<-ticker
		if !f.current.allowedDown() {
			if f.isGameOver() {
				fmt.Println("Game over :(")
				break
			}

			f.newCurrent()
		} else {
			f.current.moveDown()
		}

		f.render()

	}

}
