package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type Point struct {
	X, Y, DX, DY int
}

func main() {
	data, _ := ioutil.ReadFile("./data.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	paths := make([]Point, len(lines))

	for i, line := range lines {
		p := Point{}
		fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &p.X, &p.Y, &p.DX, &p.DY)
		paths[i] = p
	}

	min := 999999.9
	k := 0
	for {
		step(1, paths)
		value := getValue(paths)
		if value < min {
			min = value
		} else {
			step(-1, paths)
			testx, testy, basex, basey := getMin(paths)
			paint(k, paths, testx, testy, basex, basey)
			break
		}

		k++
	}

	fmt.Printf("Seconds: %v\n", k)
}
func step(ind int, paths []Point) {
	for i := range paths {
		paths[i].X += paths[i].DX * ind
		paths[i].Y += paths[i].DY * ind
	}
}
func paint(k int, paths []Point, x, y, bx, by int) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{x + 1, y + 1}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{255, 0, 0, 0xff}

	// Set color for each pixel.
	for _, p := range paths {
		img.Set(p.X-bx, p.Y-by, cyan)
		//fmt.Printf("%d %d\n", p.X, p.Y)
	}

	// Encode as PNG.
	fname := fmt.Sprintf("%d-message.png", k)
	f, _ := os.Create(fname)
	png.Encode(f, img)
}

func getMiddle(paths []Point) (int, int) {
	var mx, my int

	for _, p := range paths {
		mx += p.X
		my += p.Y
	}

	return (mx / len(paths)), (my / len(paths))
}
func getValue(paths []Point) float64 {
	_, y := getMiddle(paths)
	v := 0.0

	for _, p := range paths {
		diff := float64(p.Y - y)
		v += diff * diff
	}

	return math.Sqrt(v / float64(len(paths)))
}
func getMin(paths []Point) (int, int, int, int) {
	var mnx, mxx int
	var mny, mxy int

	for _, p := range paths {
		if p.X > mxx {
			mxx = p.X
		}
		if p.X < mnx {
			mnx = p.X
		}
		if p.Y > mxy {
			mxy = p.Y
		}
		if p.Y < mny {
			mny = p.Y
		}
	}

	mxx = mxx - mnx
	mxy = mxy - mny

	return mxx, mxy, mnx, mny
}
