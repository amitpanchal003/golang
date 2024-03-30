// looping statement
package main

import "fmt"

func main() {

	fmt.Println("welcome to day_08 looping statement")

	//days := []string{"oneday", "tuesday", "whenwhatday", "thridDay", "fouthDay", "saturday", "sunday"}
	//printing list using for...loop

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//another way of for loop
	// its like for in range syntax in python

	// for i := range days { // in this case i provide a index of value.
	// 	fmt.Println(days[i]) //using that index number it print the value present on that index number
	// }
	//fmt.Println("total values in day list is: ", len(days))
	// it print the total number of element inside the list

	// another way
	// its a kind of for each in javascript
	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	//while loop like syntax in golang
	Value := 1

	for Value < 10 {
		fmt.Println("value is: ", Value)
		Value++
	}

}
