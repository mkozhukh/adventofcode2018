package main

import (
	"fmt"
	"strconv"
)

type Step struct {
	Name     byte
	Ready    int
	Requires []*Step
	Triggers []*Step
}

func main() {
	board := make([]byte, 100000000)
	size := 2
	board[0] = 3
	board[1] = 7

	elves := make([]int, 2)
	elves[0] = 0
	elves[1] = 1

	target := 327901 + 10
	for target > size {
		count := board[elves[0]] + board[elves[1]]
		if count > 9 {
			board[size] = count / 10
			size++
			board[size] = count % 10
			size++
		} else {
			board[size] = count
			size++
		}

		for i := range elves {
			elves[i] = (elves[i] + int(board[elves[i]]) + 1) % size
		}
	}

	result := ""
	for i := target - 10; i < target; i++ {
		result += strconv.Itoa(int(board[i]))
	}
	fmt.Printf("Last 10: %s\n", result)

	size = 2
	elves[0] = 0
	elves[1] = 1
	compare := []byte{3, 2, 7, 9, 0, 1}
	for {
		count := board[elves[0]] + board[elves[1]]
		if count > 9 {
			board[size] = count / 10
			size++
			board[size] = count % 10
			size++

			if match(board, compare, size-2) || match(board, compare, size-1) {
				break
			}
		} else {
			board[size] = count
			size++

			if match(board, compare, size-1) {
				break
			}
		}

		for i := range elves {
			elves[i] = (elves[i] + int(board[elves[i]]) + 1) % size
		}
	}
}

func match(board []byte, compare []byte, size int) bool {
	if size < len(compare) {
		return false
	}

	for i := range compare {
		if board[size-5+i] != compare[i] {
			return false
		}
	}

	fmt.Printf("Found at: %d\n", size-5)
	return true
}
