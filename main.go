package main

import (
	"math/rand"
	"time"
)

const tSize = 4

func main() {
	rand.Seed(time.Now().UnixNano())

	f := Field{}
	f.init().render().placeCurrent()

	ticker := time.Tick(time.Second)

	for true {
		<-ticker
		f.removeCurrent()
		f.current.moveDown()
		f.placeCurrent()

		f.render()

	}

}
