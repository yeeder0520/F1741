package main

import (
	"errors"
	"fmt"
	"os"
)

func checkFile(filename string) {
	finfo, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("%v: 檔案不存在!\n\n", filename)
			return
		}
	}
	fmt.Printf("檔名: %s\n是目錄: %t\n修改時間: %v\n權限: %v\n大小: %d\n\n",
		finfo.Name(), finfo.IsDir(), finfo.ModTime(), finfo.Mode(), finfo.Size())
}

func main() {
	checkFile("junk.txt")
	checkFile("test.txt")
}
