package main

import (
	"fmt"
)

func newNode(value int, prev, next *Node) *Node {
	node := Node{}
	node.Value = value

	if prev == nil {
		node.Prev = &node
	} else {
		node.Prev = prev
		prev.Next = &node
	}

	if next == nil {
		node.Next = &node
	} else {
		node.Next = next
		next.Prev = &node
	}

	return &node
}

func removeNode(node *Node) *Node {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	return node.Next
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func main() {
	curr := newNode(0, nil, nil)
	players := make([]int, 418)
	maxbid := 70769 * 100

	bid := 1
	for {
		i := (bid - 1) % len(players)
		if bid%23 == 0 {
			curr = curr.Prev.Prev.Prev.Prev.Prev.Prev.Prev
			players[i] += bid + curr.Value
			curr = removeNode(curr)
		} else {
			curr = newNode(bid, curr.Next, curr.Next.Next)
		}

		if bid == maxbid {
			break
		}
		bid++
	}

	max := 0
	for i := range players {
		if players[i] > max {
			max = players[i]
		}
	}
	fmt.Printf("winner %d\n", max)
}

func log(node *Node) {
	for i := 0; i < 25; i++ {
		fmt.Printf("%d, ", node.Value)
		node = node.Next
	}
	fmt.Print("\n")
}
