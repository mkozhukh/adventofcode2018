package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type LineData struct {
	Step int
}
type WorldData struct {
	SolutionA int
	SolutionB int

	Frequency      int
	SolutionBReady bool
	Once           map[int]bool
	Data           []*LineData
}

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	// prepare the world
	world := WorldData{
		Frequency: 0,
		Once:      make(map[int]bool),
		Data:      make([]*LineData, len(lines)),
	}

	// parse incoming data
	for i, line := range lines {
		step, _ := strconv.Atoi(line)
		world.Data[i] = &LineData{step}
	}

	world.solve()
	fmt.Printf("Result A: %d\n", world.SolutionA)
	fmt.Printf("Result B: %d\n", world.SolutionB)
}

func (w *WorldData) solve() {
	w.doMath()
	w.SolutionA = w.Frequency

	for !w.SolutionBReady {
		w.doMath()
	}
}

func (w *WorldData) doMath() {
	for i := range w.Data {
		w.Frequency += w.Data[i].Step
		if !w.SolutionBReady {
			if _, ok := w.Once[w.Frequency]; ok {
				w.SolutionB = w.Frequency
				w.SolutionBReady = true
				return
			} else {
				w.Once[w.Frequency] = true
			}
		}
	}
}
