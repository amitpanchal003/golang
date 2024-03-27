
/*
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
*/

/*
	MemoryManagement
	methods.
		1. new()
			- in this memory is allocated but not initialised.
			- yes, ofcourse you'll get memory address.
			- zeroed storage [can not store any data directly]

		2. make()
			- in ths memory allocated and also initialised.
			- yes, ofcourse you'll get memory address.
			- non-zeroed storage [you can store data]
		3. garbage collection.
			- it happen automatically
			- anything that is outof the scope and nil



	Pointers.
		creating pointer
		var prt *int // defined ponter type is integer
		fmt.println(ptr)
		// output is <nil>
		// bcoz defaul value of pointer is nil


		myNumber := 25

		var ptr = &myNumber // creating referance-ponter
		// with this line we are just taking the referance of myNumber
		// & is used for referancing purpose
		// * is used for direct declaration of pointer

		fmt.println(ptr)// this will pring the memory address of myNumber
		fmt.println(*ptr)// it"ll prit actual value present in that memory

*/
