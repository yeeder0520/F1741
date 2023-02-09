package main

import "fmt"

func main() {
	v := 1234.678915
	fmt.Printf("%10.0f\n", v)
	fmt.Printf("%10.1f\n", v)
	fmt.Printf("%10.2f\n", v)
	fmt.Printf("%10.3f\n", v)
	fmt.Printf("%10.4f\n", v)
	fmt.Printf("%10.5f\n", v)
}
