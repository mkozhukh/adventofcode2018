package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	ID   int
	Kids []*Node
	Data []int
}

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), " ")
	fmt.Printf("Input parsed, %d records\n", len(lines))

	vroot := Node{Kids: make([]*Node, 0, 1)}
	pos := makeNode(&lines, 0, &vroot)
	root := vroot.Kids[0]

	fmt.Printf("parsed %d, kids %d, meta %d\n", pos, len(root.Kids), len(root.Data))
	fmt.Printf("summ %d\n", summMeta(root))
	fmt.Printf("key %d\n", keyMeta(root))
}

func makeNode(lines *[]string, pos int, node *Node) int {
	kids, _ := strconv.Atoi((*lines)[pos])
	pos++
	meta, _ := strconv.Atoi((*lines)[pos])
	pos++

	root := Node{}
	root.Kids = make([]*Node, 0, kids)
	root.Data = make([]int, meta)

	node.Kids = append(node.Kids, &root)

	for i := 0; i < kids; i++ {
		pos = makeNode(lines, pos, &root)
	}

	for i := 0; i < meta; i++ {
		root.Data[i], _ = strconv.Atoi((*lines)[pos])
		pos++
	}

	return pos
}

func summMeta(node *Node) int {
	summ := 0
	for _, kid := range node.Kids {
		summ += summMeta(kid)
	}
	for _, a := range node.Data {
		summ += a
	}

	return summ
}

func keyMeta(node *Node) int {
	summ := 0
	if len(node.Kids) == 0 {
		for _, a := range node.Data {
			summ += a
		}
	} else {
		for _, a := range node.Data {
			if a == 0 {
				continue
			}
			if a > len(node.Kids) {
				continue
			}

			summ += keyMeta(node.Kids[a-1])
		}
	}

	return summ
}
