// comma ok syntax. and user input
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin) // creating listener
	// bufio is package for reading the user input
	// it require os package .
	// basically we are saying that hey, bufio read the user input from os standard input.
	fmt.Println("enter the rating for our pizza") // adding instruction to the user

	// Comma ok || error ok syntax

	input, _ := reader.ReadString('\n') // storing user input in input variable
	// \n is for read the line until new line appear.
	// reader.ReadString will read only string
	fmt.Println("thans for rating", input)    // printing the rating
	fmt.Printf("type of rating is %T", input) // %T checking the type of rating
}
