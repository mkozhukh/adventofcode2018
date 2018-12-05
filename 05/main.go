package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	fmt.Printf("Input parsed, %d bytes\n", len(data))

	fmt.Printf("Final polymer length: %d\n", runReaction(data, 0x00))

	bases := []byte("QWERTYUIOPASDFGHJKLZXCVBNM")
	minLen := len(data) + 1
	for _, ex := range bases {
		exLen := runReaction(data, rune(ex))
		if exLen < minLen {
			minLen = exLen
		}
	}

	fmt.Printf("Minimal polymer length: %d\n", minLen)
}

func runReaction(data []byte, exclude rune) int {
	result := make([]byte, len(data))
	i := 0
	for _, ru := range data {
		if exclude != 0x00 && unicode.ToUpper(rune(ru)) == exclude {
			continue
		}

		if i > 0 {
			lu := result[i-1]
			if lu != ru && unicode.ToUpper(rune(lu)) == unicode.ToUpper(rune(ru)) {
				i--
				continue
			}
		}

		result[i] = ru
		i++
	}

	return i
}
