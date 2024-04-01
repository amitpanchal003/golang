package main

import (
	"fmt"
	"io"
	"os"
)

// solve this errors
func main() {
	Files()

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

func Files() {
	content := "this demo for writing files in golang"

	file, err := os.Create("./myfile.txt")
	// while creating file there might be errors, we have to handle this

	// if err != nil {
	// 	panic(err) // panic will just shuttdown the execution of program and return the err

	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	// using io package, pushing content into file

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	// why defer, bcoz we want close file after everything is done

	readFile("./myfile.txt")

}

func readFile(filename string) {
	//for reading the file there is seperate package which is [ioutil]

	dataByte, err := os.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("text data inside the file is: \n", string(dataByte))
}

//short hand for checking error

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
