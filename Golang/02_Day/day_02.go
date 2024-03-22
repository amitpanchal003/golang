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
package main

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
