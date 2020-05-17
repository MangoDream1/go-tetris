package main

func fixShapeLeft(shape *[tSize][tSize]bool) {
	maxLeft := tSize

	for _, row := range shape {
		for i, value := range row {
			if value {
				if i < maxLeft {
					maxLeft = i
					continue
				}
			}
		}
	}

	if maxLeft == 0 {
		return
	}

	for iRow, row := range shape {
		for i, value := range row {
			if value {
				shape[iRow][i] = false
				shape[iRow][i-maxLeft] = true
			}
		}
	}
}

func fixShapeUp(shape *[tSize][tSize]bool) {
	maxUp := tSize

	for iRow, row := range shape {
		for _, value := range row {
			if value {
				if iRow < maxUp {
					maxUp = iRow
					break
				}
			}
		}
	}

	if maxUp == 0 {
		return
	}

	for iRow, row := range shape {
		for i, value := range row {
			if value {
				shape[iRow][i] = false
				shape[iRow-maxUp][i] = true
			}
		}
	}
}
