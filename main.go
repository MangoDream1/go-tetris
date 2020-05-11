package main

import (
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
			f.setCurrent()
		}

		f.current.moveDown()

		f.render()

	}

}
