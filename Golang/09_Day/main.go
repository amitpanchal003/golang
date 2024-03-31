//functions in golang

package main

import "fmt"

func add(x int, y int) int {
	// it is important to mention which type of data function will return.
	return x + y
}
func Greetings() { // normal function declaration
	fmt.Println("hello there")
}
func addition(values ...int) int { //...int this is spread operator
	total := 0 //for storing the total of values

	for _, val := range values {
		total += val // this will go through the values and add that in to total
	}
	return total // then return it
}
func main() {
	result := add(10, 20)                   //calling function and storing response in result variable
	fmt.Println("addition is: ", result)    //printing the result
	Greetings()                             //function calling
	proAdd := addition(2, 3, 8, 4, 5, 5, 6) //storing result in proAdd variable
	fmt.Println("total is: ", proAdd)       // printing result
}
