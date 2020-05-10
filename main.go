package main

import (
	"time"
)

const tSize = 4

func main() {
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
