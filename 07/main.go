package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Step struct {
	Name     byte
	Ready    int
	Requires []*Step
	Triggers []*Step
}

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	paths := make(map[byte]*Step, len(lines)*2)

	for _, line := range lines {
		a := []byte(line)[5]
		b := []byte(line)[36]

		before, ok := paths[a]
		if !ok {
			before = &Step{a, 0, make([]*Step, 0), make([]*Step, 0)}
			paths[a] = before
		}

		after, ok := paths[b]
		if !ok {
			after = &Step{b, 0, make([]*Step, 0), make([]*Step, 0)}
			paths[b] = after
		}

		after.Requires = append(after.Requires, before)
		before.Triggers = append(before.Triggers, after)
		after.Ready++
	}

	receipt := make([]byte, 0, len(lines))
	for {
		var found *Step
		min := []byte("Z")[0]

		for _, test := range paths {
			if test.Ready == 0 && test.Name <= min {
				min = test.Name
				found = test
			}
		}

		if found == nil {
			break
		}

		receipt = append(receipt, found.Name)

		found.Ready--
		for _, next := range found.Triggers {
			next.Ready--
		}
	}

	fmt.Printf("Receipt: %s\n", receipt)

	for _, next := range paths {
		next.Ready = len(next.Requires)
	}

	time := -1
	workers := make([]int, 5)
	jobs := make([]*Step, 5)
	base := []byte("A")[0] - 1
	for {
		count := 0
		for i := range workers {
			if workers[i] != 0 {
				workers[i]--
				if workers[i] == 0 {
					for _, task := range jobs[i].Triggers {
						task.Ready--
					}
				}
			}
			if workers[i] == 0 {
				count++
			}
		}

		time++

		// all busy
		if count == 0 {
			continue
		}

		next := make([]*Step, 0)
		for _, test := range paths {
			if test.Ready == 0 {
				next = append(next, test)
			}
		}

		if len(next) == 0 {
			if count == 5 {
				break
			} else {
				continue
			}
		}

		sort.SliceStable(next, func(i, j int) bool {
			return next[i].Name < next[j].Name
		})

		j := 0
		for i := range workers {
			if workers[i] == 0 {
				workers[i] = int(next[j].Name-base) + 60
				jobs[i] = next[j]
				j++

				jobs[i].Ready--
				fmt.Printf("[%d] Worker %d gets job %s, %d secs \n", time, i, string(jobs[i].Name), workers[i])
			}
			if j == len(next) {
				break
			}
		}
	}

	fmt.Printf("Time: %d\n", time)
}
