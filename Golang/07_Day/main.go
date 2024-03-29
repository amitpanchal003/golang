// conditional statement in golang.
// 01.if statement.
// syntax
package main

import (
	"fmt"
)

func main() {
	fmt.Println("learnig conditional statement in golang")

	// amitHasMoney := true

	// if amitHasMoney == true {
	// 	fmt.Println("eligible to live in society")

	// }

	//ex:2

	/*grade := 70 //initialization of int type variable

	if grade >= 65 { // if this condition become true then only it execute inner block
		fmt.Println("passing marks") //print the statement

	}
	*/

	//ex:1 if.....elseif....else
	/*90 or above is equivalent to an A grade
	80-89 is equivalent to a B grade
	70-79 is equivalent to a C grade
	65-69 is equivalent to a D grade
	64 or below is equivalent to an F grade
	*/

	grade := 60

	if grade >= 90 {
		fmt.Println("A grade")
	} else if grade >= 80 {
		fmt.Println("B grade")
	} else if grade >= 70 {
		fmt.Println("C grade")
	} else if grade >= 65 {
		fmt.Println("D grade")
	} else {
		fmt.Println("Failing grade")
	}
}
