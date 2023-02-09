package main

import (
	"fmt"

	"Example09.04/shape"
)

func main() {
	triangle := shape.Triangle{Base: 15.5, Height: 20.1}
	rectangle := shape.Rectangle{Length: 20, Width: 10}
	square := shape.Square{Side: 10}
	shapes := [...]shape.Shape{triangle, rectangle, square}

	for _, item := range shapes {
		fmt.Printf("%s的面積: %.2f\n", shape.GetName(item), shape.GetArea(item))
	}
}
