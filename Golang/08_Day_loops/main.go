// looping statement
package main

import "fmt"

func main() {
	fmt.Println("welcome to day_08 looping statement")

	days := []string{"oneday", "tuesday", "whenwhatday", "thridDay", "fouthDay", "saturday", "sunday"}
	//printing list using for...loop

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

}
