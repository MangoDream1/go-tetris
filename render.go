package main

import (
	"fmt"
)

const vertPadding = 2
const hortPadding = 1

func findMax(arguments ...int) int {
	result := 0
	for _, a := range arguments {
		if a > result {
			result = a
		}
	}

	return result
}

func (f *Field) render() *Field {
	field := _renderField(f)
	coming := _renderComing(f)
	stored := _renderStored(f)

	output := make([][]byte, findMax(len(field), len(coming), len(stored)))

	sections := [][][]byte{stored, field, coming}

	prevSectionsLength := 0
	for _, section := range sections {
		for i := range section {
			if len(output[i]) > 0 {
				output[i] = append(output[i], section[i]...)
			} else {
				output[i] = append(make([]byte, prevSectionsLength), section[i]...)
			}
		}

		prevSectionsLength += len(section[0])
	}

	// render empty bytes as space
	for i, row := range output {
		for j, b := range row {
			if b == byte(0) {
				output[i][j] = ' '
			}
		}
	}

	fmt.Printf("\033[2;0H") // move cursor to second row and draw from there
	for _, row := range output {
		fmt.Println(string(row))
	}

	return f
}

func _renderStored(f *Field) [][]byte {
	output := make([][]byte, tSize)

	for i := 0; i < tSize; i++ {
		output[i] = make([]byte, tSize)
	}

	if f.stored == (Tetromino{}) {
		return output
	}

	for i, row := range f.stored.shape {
		for j, value := range row {
			if value {
				output[i][j] = f.stored.value
			}
		}
	}

	return output
}

func _renderComing(f *Field) [][]byte {
	output := make([][]byte, maxStored*(tSize+hortPadding))

	offset := 0
	for _, t := range f.coming {
		shapeHeight := 0

		for i, row := range t.shape {
			output[i+offset] = make([]byte, tSize)

			counted := false
			for j, value := range row {
				if value {
					output[i+offset][j+1] = t.value

					if !counted {
						shapeHeight++
						counted = true
					}
				}
			}
		}

		offset += shapeHeight
		shapeHeight = 0

		for paddingI := 0; paddingI < hortPadding; paddingI++ {
			output[offset+paddingI] = make([]byte, tSize)
			offset++
		}
	}

	return output
}

func _renderField(f *Field) [][]byte {
	output := make([][]byte, fieldHeight+1)

	for i, row := range f.state {
		output[i] = make([]byte, fieldWidth+2)

		// define edges
		output[i][0] = '|'
		output[i][fieldWidth+1] = '|'

		for j, b := range row {
			output[i][j+1] = b
		}
	}

	// add bottom
	output[fieldHeight] = make([]byte, fieldWidth+2)
	for i := 0; i < fieldWidth+2; i++ {
		output[fieldHeight][i] = '='
	}

	return output
}
