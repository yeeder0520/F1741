package main

import (
	"fmt"
)

func main() {
	itemsSold()
}

func itemsSold() {
	items := make(map[string]int)
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24

	for k, v := range items {
		fmt.Printf("%s 賣出 %d 件商品, 表現", k, v)
		if v < 40 {
			fmt.Println("低於預期.")
		} else if v > 40 && v <= 100 {
			fmt.Println("符合預期")
		} else if v > 100 {
			fmt.Println("超乎預期")
		}
	}
}
