package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Star struct {
	X, Y     int
	Area     int
	IsFinite bool
}

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	size := 400
	stars := make([]Star, len(lines))

	for i, line := range lines {
		star := Star{IsFinite: true, Area: 0}
		fmt.Sscanf(line, "%d, %d", &star.X, &star.Y)
		stars[i] = star
	}

	for i := -1; i <= size; i++ {
		for j := -1; j <= size; j++ {
			assignArea(&stars, i, j, size)
		}
	}

	res := Star{Area: 0}
	for _, star := range stars {
		if star.IsFinite && star.Area > res.Area {
			res = star
		}
	}

	fmt.Printf("Largest finite area: %d\n", res.Area)

	fmt.Printf("Safe area: %d\n", getSafeArea(stars))
}

func assignArea(stars *[]Star, x, y, size int) {
	var min float64
	var res int

	none := false
	min = float64(size) * 3.0
	for i, star := range *stars {
		dist := math.Abs(float64(x-star.X)) + math.Abs(float64(y-star.Y))
		if dist < min {
			min = dist
			res = i
		} else if dist == min {
			none = true
		}
	}

	if !none {
		if x < 0 || y < 0 || x >= size || y >= size {
			(*stars)[res].IsFinite = false
		} else {
			(*stars)[res].Area++
		}
	}

}

func getSafeArea(stars []Star) int {
	area := 0

	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			sum := 0.0
			for _, star := range stars {
				sum += math.Abs(float64(i-star.X)) + math.Abs(float64(j-star.Y))
			}
			if sum < 10000 {
				area++
			}
		}
	}

	return area
}
