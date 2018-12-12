package main

import (
	"fmt"
	"io/ioutil"
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

	//zero := 25
	field := "........................." + lines[0] + "........................."
	full := len(field)

	bases := lines[2:]
	grows := make([]string, 0, len(lines)-2)
	for _, match := range bases {
		text := []rune(match)
		if text[9] == []rune("#")[0] {
			grows = append(grows, string(text[0:5]))
		}
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("%s\n", field)
		nextAr := []byte(field)
		for j := 2; j < full-2; j++ {
			for _, pattern := range grows {
				if field[j-2:j+3] == pattern {
					nextAr[j] = []byte("#")[0]
					break
				}
				nextAr[j] = []byte(".")[0]
			}
		}
		field = string(nextAr)
	}

	values := strings.Split(field, "")
	count := 0
	for i, test := range values {
		if test == "#" {
			count += (i - 25)
		}
	}

	fmt.Printf("summ: %d\n", count)

	second := []byte("........................." + lines[0] + ".....")
	doFast(second, 100)
}

func doFast(input []byte, count int) {
	set := make([]byte, 10000000)
	var i uint
	for i = 0; i < uint(len(input)-5); i++ {
		var j uint
		for j = 0; j < 5; j++ {
			if input[i+j] == []byte("#")[0] {
				set[i+2] += byte(1 << j)
			}
		}
	}
	print(0, 150, set)

	zero := 0.0
	min := 0
	max := 150
	for i := 0; i < count; i++ {
		var next byte
		next = 0
		nm := len(set)
		nx := 0
		for j := min; j < max; j++ {
			test := set[j]
			if test == 6 || test == 15 || test == 19 || test == 11 || test == 24 || test == 27 || test == 4 ||
				test == 18 || test == 26 || test == 2 || test == 23 || test == 9 || test == 30 || test == 29 ||
				test == 22 || test == 7 || test == 21 || test == 28 {

				set[j] = 4 + next
				set[j-1] += 8
				set[j-2] += 16
				next = (next >> 1) + 2

				if nm > j-3 {
					nm = j - 3
				}
				if nx < j+3 {
					nx = j + 3
				}
			} else {
				set[j] = next
				next = (next >> 1)
			}
		}
		min = nm
		max = nx

		if min > 750000 {
			for i := min; i <= max+1; i++ {
				set[i-500000] = set[i]
				set[i] = 0
			}

			min -= 500000
			max -= 500000
			zero += 500000
		}
		//print(0, 150, set)
	}

	fmt.Printf("After %d, we have %v (range %d - %d)\n", count, sum(25, zero, set), min, max)
}

func print(start int, end int, data []byte) {
	for i := start; i < end; i++ {
		//		fmt.Printf("%d, ", data[i])
		if data[i]&4 == 4 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Print("\n")
}

func sum(start int, zero float64, data []byte) float64 {
	summ := 0.0
	for i := start; i < len(data); i++ {
		if data[i]&4 == 4 {
			summ += float64(i-start) + zero
		}
	}

	return summ
}
