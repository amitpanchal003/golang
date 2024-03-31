//functions in golang

package main

import "fmt"

func add(x int, y int) int {
	// it is important to mention which type of data function will return.
	return x + y
}

func main() {
	result := add(10, 20)                //calling function and storing response in result variable
	fmt.Println("addition is: ", result) //printing the result
}
