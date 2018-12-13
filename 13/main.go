package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const (
	Right  = 0
	Bottom = 1
	Left   = 2
	Top    = 3
)

type Cart struct {
	X, Y, Dir, State int
	Crashed          bool
}

func main() {
	chars := []byte("/\\+")
	TurnRight := chars[0]
	TurnLeft := chars[1]
	Intersection := chars[2]

	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), "\n")

	height := len(lines)
	width := len(lines[0])

	tracks := make([][]byte, height)
	for i := range lines {
		tracks[i] = []byte(lines[i])
	}

	carts := make([]*Cart, 0)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			test := tracks[i][j]
			if string(test) == ">" {
				carts = append(carts, &Cart{j, i, Right, 0, false})
				tracks[i][j] = []byte("-")[0]
			}
			if string(test) == "<" {
				carts = append(carts, &Cart{j, i, Left, 0, false})
				tracks[i][j] = []byte("-")[0]
			}
			if string(test) == "^" {
				carts = append(carts, &Cart{j, i, Top, 0, false})
				tracks[i][j] = []byte("|")[0]
			}
			if string(test) == "v" {
				carts = append(carts, &Cart{j, i, Bottom, 0, false})
				tracks[i][j] = []byte("|")[0]
			}
		}
	}

	var crashed *Cart
	var crashedTime int
	tick := 0

	for {
		deleted := 0
		for i := range carts {
			if carts[i-deleted].Crashed {
				carts[i-deleted] = carts[len(carts)-deleted-1]
				deleted++
			}
		}
		if deleted > 0 {
			carts = carts[:len(carts)-deleted]
		}

		if len(carts) == 1 {
			break
		}

		sort.SliceStable(carts, func(i, j int) bool {
			if carts[i].Y == carts[j].Y {
				return carts[i].X < carts[j].X
			}
			return carts[i].Y < carts[j].Y
		})

		for _, cart := range carts {
			switch cart.Dir {
			case Top:
				cart.Y--
			case Bottom:
				cart.Y++
			case Left:
				cart.X--
			case Right:
				cart.X++
			}

			for _, check := range carts {
				if check != cart && !check.Crashed && check.X == cart.X && check.Y == cart.Y {
					if crashed == nil {
						crashed = cart
						crashedTime = tick
					}

					cart.Crashed = true
					check.Crashed = true
					break
				}
			}

			switch tracks[cart.Y][cart.X] {
			case TurnRight:
				switch cart.Dir {
				case Top:
					cart.Dir = Right
				case Bottom:
					cart.Dir = Left
				case Right:
					cart.Dir = Top
				case Left:
					cart.Dir = Bottom
				}

			case TurnLeft:
				switch cart.Dir {
				case Top:
					cart.Dir = Left
				case Bottom:
					cart.Dir = Right
				case Right:
					cart.Dir = Bottom
				case Left:
					cart.Dir = Top
				}

			case Intersection:
				switch cart.State % 3 {
				case 0:
					cart.Dir = (cart.Dir + 4 - 1) % 4
				case 2:
					cart.Dir = (cart.Dir + 1) % 4
				}
				cart.State++
			}
			//fmt.Printf("[%d] %v\n", i, *cart)
		}
		tick++
	}

	fmt.Printf("%d carts detected\n", len(carts))
	fmt.Printf("crash at %d,%d, after %d secs\n", crashed.X, crashed.Y, crashedTime)
	fmt.Printf("last at %d,%d, after %d secs\n", carts[0].X, carts[0].Y, tick)
}
