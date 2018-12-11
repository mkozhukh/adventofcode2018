package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y, DX, DY int
}

func main() {
	cells := make([]float64, 301*301)

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			rack := x + 10
			power := (rack*y + 3628) * rack
			cells[x+y*301] = math.Floor(float64((power%1000)/100)) - 5
		}
	}

	summ, dx, dy := getMax(cells, 3)
	fmt.Printf("Max value: %d, at %d,%d\n", int(summ), dx, dy)

	size := 3
	for i := 4; i <= 300; i++ {
		temp, x, y := getMax(cells, i)
		if temp > summ {
			size = i
			dx = x
			dy = y
			summ = temp
		}
	}

	// too lazy to implement
	// https://en.wikipedia.org/wiki/Summed-area_table
	fmt.Printf("Max value: %d, at %d,%d,%d\n", int(summ), dx, dy, size)
}

func getMax(cells []float64, size int) (float64, int, int) {
	dx := 1
	dy := 1
	summ := 0.0
	for x := 1; x <= 300-size; x++ {
		for y := 1; y <= 300-size; y++ {
			temp := 0.0
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					temp += cells[x+i+(y+j)*301]
				}
			}

			if temp > summ {
				dx = x
				dy = y
				summ = temp
			}
		}
	}

	return summ, dx, dy
}
