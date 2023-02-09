package main

import (
	"fmt"
	"time"
)

func displayTimeZone(t time.Time) {
	fmt.Print("Time: ", t, "\nTimezone: ", t.Location(), "\n\n")
}

func main() {
	date := time.Date(2021, 4, 22, 16, 44, 05, 324359102, time.Local)
	timeZone1, _ := time.LoadLocation("America/New_York")
	remoteTime1 := date.In(timeZone1)
	timeZone2, _ := time.LoadLocation("Australia/Sydney")
	remoteTime2 := date.In(timeZone2)
	timeZone3 := time.FixedZone("My TimeZone", -1*60*60)
	remoteTime3 := date.In(timeZone3)

	displayTimeZone(date)
	displayTimeZone(remoteTime1)
	displayTimeZone(remoteTime2)
	displayTimeZone(remoteTime3)
}
