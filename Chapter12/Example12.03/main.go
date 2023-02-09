package main

import "os"

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("使用 Write() 寫入\n"))
	f.WriteString("使用 Writestring() 寫入\n")
}
