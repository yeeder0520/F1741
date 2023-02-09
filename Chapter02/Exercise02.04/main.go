package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("輸入值不得為負")
	} else if input > 100 {
		return errors.New("輸入值不得超過 100")
	} else if input%7 == 0 {
		return errors.New("輸入值不得為 7 的倍數")
	} else {
		return nil
	}
}

func main() {
	input := 21
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "是偶數")
	} else {
		fmt.Println(input, "是奇數")
	}
}
