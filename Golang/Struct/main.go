// learing about struc

package main

import "fmt"

func main() {
	fmt.Println("exploring struct syntax")
	//there is inharitance in golang also no parent and child, no super

	//creating user; it'll require the value in defined structure manner
	// like Name,Email,Status,status
	creatUser := User{"amit", "amit@gmail.com", true, 28}

	fmt.Println("user details is: ", creatUser)
	fmt.Printf("user details is: %+v\n", creatUser)

}

// defing struct

type User struct {
	Name   string
	Emai   string
	Status bool
	Age    int
}
