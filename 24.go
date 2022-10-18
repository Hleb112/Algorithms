package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func newPoint(x int, y int) Point {
	return Point{x: x, y: y}
}

func main() {
	point := newPoint(0, 0)
	point2 := newPoint(5, 5)
	fmt.Println(point)
	fmt.Println(point2)
	f_dl := math.Sqrt(math.Pow(float64(point2.x-point.x), 2) + math.Pow(float64(point2.y-point.y), 2))
	fmt.Println(f_dl)
}
