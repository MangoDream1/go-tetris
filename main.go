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

	ticker := time.Tick(time.Second)

	f.storeCurrent()

	for true {
		<-ticker
		f.current.remove()
		f.current.moveDown()
		f.current.place()

		f.render()

	}

}
