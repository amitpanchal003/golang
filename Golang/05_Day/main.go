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
}
