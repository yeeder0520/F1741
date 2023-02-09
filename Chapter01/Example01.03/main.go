package main

import "math/rand"

func main() {
	var seed = 1234456789
	rand.Seed(seed) // 正確寫法: var seed int64 = 1234456789
}
