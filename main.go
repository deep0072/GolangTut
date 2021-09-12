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

	// if i want week day with time then pass "Monday" "15:04:05" respectively
	fmt.Println(PresentTime.Format("01-02-2006  15:040:05 Monday"))
}
