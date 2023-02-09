package main

import "fmt"

type name string

type point struct {
	x int
	y int
}

func (n name) printName() {
	fmt.Println("name:", n)
}

func (p *point) setPoint(x, y int) {
	p.x = x
	p.y = y
}

func (p point) getPoint() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func main() {
	var n name = "Golang"
	n.printName()

	a, b := point{}, point{}
	a.setPoint(10, 10)
	b.setPoint(10, 5)
	fmt.Println("point1:", a.getPoint())
	fmt.Println("point2:", b.getPoint())
}
