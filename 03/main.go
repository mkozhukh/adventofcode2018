package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

type Claim struct {
	image.Rectangle
	ID int
}

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	// prepare the world
	claims := make([]Claim, len(lines))
	for i, line := range lines {
		var id, x, y, width, height int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)
		claims[i] = Claim{
			image.Rectangle{image.Point{x, y}, image.Point{x + width, y + height}},
			id,
		}
	}

	world := make([]int, 1000*1000)
	for _, patch := range claims {
		for x := patch.Min.X; x < patch.Max.X; x++ {
			for y := patch.Min.Y; y < patch.Max.Y; y++ {
				world[1000*y+x]++
			}
		}
	}

	summ := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if world[1000*y+x] > 1 {
				summ++
			}
		}
	}

	fmt.Printf("Result A: %d\n", summ)

	for _, base := range claims {
		check := true
		for _, compare := range claims {
			if base.ID != compare.ID && base.Overlaps(compare.Rectangle) {
				check = false
				break
			}
		}

		if check {
			fmt.Printf("Result B: %d\n", base.ID)
			break
		}
	}
}
