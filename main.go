package main

import (
	"fmt"
)

const tSize = 4

func main() {
	t := Tetromino{shape2.shape, shape2.value, 0, 0}

	t.moveDown()

	fmt.Println(t)

}
