package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Name() string
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	s := square{side: 10}
	c := circle{radius: 6.4}
	t := triangle{base: 15.5, height: 20.1}
	printShapeDetails(s, c, t)
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("%s的面積: %.2f\n", item.Name(), item.Area())
	}
}

func (c circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) Name() string {
	return "圓形"
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Name() string {
	return "正方形"
}

func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t triangle) Name() string {
	return "三角形"
}
