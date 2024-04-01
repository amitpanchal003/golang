//Defer in golang

package main

import "fmt"

func main() {
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("ONE")
	fmt.Println("welcome to day 10")

}
