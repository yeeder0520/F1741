package main

import "fmt"

func main() {
	taxTotal := .0

	taxTotal += salesTax(.99, .075)
	taxTotal += salesTax(2.75, .015)
	taxTotal += salesTax(.87, .02)
	fmt.Println("Sales Tax Total: ", taxTotal)
}

func salesTax(cost float64, taxRate float64) float64 {
	return cost * taxRate
}
