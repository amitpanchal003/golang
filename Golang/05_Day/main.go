package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// to gate date time and day in human readable format
	// we have to pass this parameter in required format

	// creating time and date
	createDate := time.Date(2020, time.November, 22, 06, 12, 0, 0, time.UTC)

	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
