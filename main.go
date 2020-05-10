package main

import (
	"time"
)

const tSize = 4

func main() {
	f := Field{}
	f.init().render()

	ticker := time.Tick(time.Second)

	for true {
		<-ticker
		f.render()
	}

}
