package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var ErrWorkingFileNotFound = errors.New("查無工作檔案")

func main() {
	workFileName := "note.txt"
	backupFileName := "backup.txt"
	err := writeBackup(workFileName, backupFileName)
	if err != nil {
		panic(err)
	}
}

func writeBackup(work, backup string) error {
	workFile, err := os.Open(work)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrWorkingFileNotFound
		}
		return err
	}
	defer workFile.Close()

	backFile, err := os.OpenFile(backup, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer backFile.Close()

	content, err := io.ReadAll(workFile)
	if err != nil {
		return err
	}

	_, err = backFile.WriteString(fmt.Sprintf("[%v]\n%v", time.Now().String(), string(content)))
	if err != nil {
		return err
	}

	return nil
}
