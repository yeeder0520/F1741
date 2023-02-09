package main

import (
	"fmt"
	"log"
	"time"
)

const formatString = "15:04:05 01/02/2006"

func deadline(t time.Time, tz string) (time.Time, error) {
	target := t.AddDate(5, 0, 0)
	timeZone, err := time.LoadLocation(tz)
	if err != nil {
		return t, fmt.Errorf("時區錯誤")
	}
	return target.In(timeZone), nil
}

func timeFormat(t time.Time) string {
	return t.Format(formatString)
}

func main() {
	now := time.Now()
	remoteTime, err := deadline(now, "Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("現在時間:", timeFormat(now))
	fmt.Println("期限    :", timeFormat(remoteTime))
}
