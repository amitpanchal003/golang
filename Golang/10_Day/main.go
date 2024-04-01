//Defer in golang

package main

import "fmt"

func main() {
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("ONE")
	fmt.Println("welcome to day 10")

	// A defer statement defers the execution of a function until the
	// surrounding function returns.

	//The deferred call's arguments are evaluated immediately,
	//but the function call is not executed until the surrounding function returns.

	// in simple term , it work like LIFO

}
