// Exploring data type and structure.
/*
	two type of Data.
	01. primitive data type.
	02. composit data type.

		01. primitive data type:
			Think of primitive data types as atoms.
			Atoms are the smallest particles in the universe
			which cannot be further broken.
			(Though we can use the latest technology to break them
			 but we will leave it for the physicists.) Like atoms,
			 primitive data types also cannot be further broken down.
			 And like atoms, on combining them, we can form complex data
			 types but more on that later.
		>> In golang, we have int, float, byte, string, rune & bool



*/
//Problem #1: We want to populate the information of a student.
/*package main

import "fmt"

func main() {
	var name string
	var age int
	var address string
	var gpa float64
	var isAdult bool

	name = "amit"
	age = 28
	address = "latur maharashtra"
	gpa = 3.9
	isAdult = true
	fmt.Println(name, age, address, gpa, isAdult)
}
*/

/*
#####################################################################
#####################################################################
Array:
An array is a systematic arrangement of similar objects,
usually in rows and columns.
>> the type[n]T is an array of n value of type T

	i.e. var a [5]int
	here,
	a = is variable name.
	[5]= is size of array.
	int = is data type

package main

import "fmt"

func main() {
	var list1 [2]string //----array declaration
	list1[0] = "hello"  // ----array initialization
	list1[1] = "amit"   // -----array initialization

	fmt.Println(list1) // -------printing the value of array

	//---------------------another way

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	//---------------------using shorthand

	primeNumbers := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primeNumbers)
}
*/
/* ###############################################
   ###############################################

>> Slices:
	_In Go language slice is more powerful, flexible,
	convenient than an array, and is a lightweight data structure.
	 Slice is a variable-length sequence that stores elements of a
	 similar type, you are not allowed to store different type of
	 elements in the same slice.


*/
package main

import "fmt"

func main() { // Array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	//struct
	type person struct {
		name string
		age  int
	}
	var p1 person = person{name: "Bob", age: 28}
	fmt.Println(p1)
	// Map
	var data1 map[string]int = map[string]int{"age": 28, "size": 180}
	fmt.Println(data1)
	fmt.Println("showing demo , how add, commit , push work")
}
