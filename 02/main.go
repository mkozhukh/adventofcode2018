package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type LineData struct {
	ID []byte
}
type WorldData struct {
	SolutionA int
	SolutionB string

	Data []*LineData
}

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	// prepare the world
	world := WorldData{
		Data: make([]*LineData, len(lines)),
	}

	// parse incoming data
	for i, line := range lines {
		world.Data[i] = &LineData{[]byte(line)}
	}

	world.solve()
	fmt.Printf("Result A: %d\n", world.SolutionA)
	fmt.Printf("Result B: %s\n", world.SolutionB)
}

func (w *WorldData) solve() {
	d := 0
	t := 0
	for _, line := range w.Data {
		dl, tl := line.letterCount()
		if dl {
			d++
		}
		if tl {
			t++
		}
	}
	w.SolutionA = d * t

	minDist := 9999
	for _, line := range w.Data {
		for _, line2 := range w.Data {
			if line != line2 {
				dist, common := line.distance(line2)
				if dist < minDist {
					minDist = dist
					w.SolutionB = common
				}
			}
		}
	}
}

func (l *LineData) distance(l2 *LineData) (int, string) {
	diff := 0
	common := make([]byte, len(l.ID))

	for i := range l.ID {
		if l.ID[i] != l2.ID[i] {
			diff++
		} else {
			common[i-diff] = l.ID[i]
		}
	}

	return diff, string(common[0 : len(l.ID)-diff])
}

func (l *LineData) letterCount() (bool, bool) {
	m := make(map[byte]int)

	for _, char := range l.ID {
		m[char]++
	}

	d := false
	t := false

	for k := range m {
		if m[k] == 2 {
			d = true
		}
		if m[k] == 3 {
			t = true
		}
	}
	return d, t
}
