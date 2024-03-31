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
func addition(values ...int) (int, string) { //...int this is spread operator
	//we can also return the string
	total := 0 //for storing the total of values

	for _, val := range values {
		total += val // this will go through the values and add that in to total
	}
	return total, "hello im making addition" // then return it
	//just add comma and in doubble quote add your string
}
func main() {
	result := add(10, 20)                      //calling function and storing response in result variable
	fmt.Println("addition is: ", result)       //printing the result
	Greetings()                                //function calling
	proAdd, _ := addition(2, 3, 8, 4, 5, 5, 6) //storing result in proAdd variable
	//using ( _ ) this we can ignore message or we can handle it like
	//proAdd, mymessage:= addition(1,2,3,4,5,6)

	//then just print it
	//fmt.Println("message is: ",mymessage)
	fmt.Println("total is: ", proAdd) // printing result
}
