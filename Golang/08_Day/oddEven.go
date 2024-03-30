// Write a program that takes a number as input and prints whether it is even or odd.
/*package main

import "fmt"

func main() {
	fmt.Println("enter your number")
	var num int
	fmt.Scanln(&num)// fmt provide scannaer class to take input fron user
	// check num is even or odd
	if num%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is even")
	}

}
*/
// using bufi package solution

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//first we have to implenent bufio reader to read input from standard input output

	reader := bufio.NewReader(os.Stdin)

	//read the line from  user
	fmt.Println("enter the number: ")
	num, err := reader.ReadString('\n')

	if err != nil { // it check is thers any error or not [empty or not]
		fmt.Println(err) //if not empty then just print that available error
		return           // return[terminate the process and return]
	}

	// after this we need to convert the input into number
	number, err := strconv.ParseInt(strings.TrimSpace(num), 36, 36)

	if err != nil {
		fmt.Println(err)
		return
	}
	// build the logic for odd-even number
	if number%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
}
