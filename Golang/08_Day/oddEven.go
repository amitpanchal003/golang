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
// using bufio package solution
/*
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
*/

// 02. Write a program that takes two numbers as input
// and prints the larger of the two numbers.
/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//creating scanner
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter first number: ")
	num, err := reader.ReadString('\n')

	num2, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}
	number, err := strconv.ParseFloat(strings.TrimSpace(num), 64)
	number2, err := strconv.ParseFloat(strings.TrimSpace(num2), 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	if number > number2 {
		fmt.Println("larger number is: ", number)
	} else if number < number2 {
		fmt.Println("larger number is: ", number2)
	}

}*/

/*
Another simple way to do this
package main

import "fmt"
func main(){
	var num1,num2 int
	fmt.println("enter first number: ")
	fmt.Scanln(&num1)
	fmt.println("enter second number: ")
	fmt.Scanln(&num2)
	if num>num2{
		"first is largest"
		else
		second is largest
	}
}
*/