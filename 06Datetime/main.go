package main

import (
	"fmt"
	"time"
)

func main() {
	PresentTime := time.Now()
	fmt.Println(PresentTime)

	// to get the time in mm-dd-yy

	// we have to pass "01-02-2006" as an string argument

	fmt.Println(PresentTime.Format("01-02-2006"))

	// if i want week day with time then pass "Monday  "15:04:05" respectively
	fmt.Println(PresentTime.Format("01-02- 2006  15:040:05 Monday"))

	// create date from manual input by using time.date(year, month, hiur, minute,second, nano sec, timezone) method

	// here time.Utc

	CreateDate := time.Date(2020, time.January, 1, 18, 18, 0, 0, time.UTC)

	fmt.Println(CreateDate)
	fmt.Println(CreateDate.Format("01-02-2006 15:04:05 Monday"))
}
