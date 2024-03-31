//Methods

package main

import "fmt"

// defining structure

type User struct {
	Name     string
	Email    string
	isActive bool
	salary   int
}

//defining method
func (u User) show() {
	fmt.Println("name is: ", u.Name)
	fmt.Println("email is: ", u.Email)
	fmt.Println("active status is: ", u.isActive)
	fmt.Println("salary is: ", u.salary)

}

func main() {
	fmt.Println("welcome to methods section")

	//initialization of user

	res := User{
		Name:     "amit",
		Email:    "test@mail.com",
		isActive: true,
		salary:   120000,
	}
	//calling the method
	res.show()
}
