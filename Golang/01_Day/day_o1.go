// hello this is day one of goLang Series.
package main

// it is a special package that is used to create executable programs.
// it is entry point for any executable program.
// it must contain functiom called main().
import "fmt"

// it provide input and output function.

func main() { // This is main method{function} this is main entry point of program.
	var age int = 28 //variable declaration and assign the value
	/*var is a keyword
	  age is varable name.
	  int is data type.
	  = is operator
	  28 is value.

	*/
	var name string = "Hi amit"
	// string type variable declaration and assignment
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("hello world") // This is output statement
}

/*
#some important things to keep in mind while declaring variables.
01. variable must br declared before they can be used.
02. the type of variable must be specified.
03. variable can de declared outside(global variable) as well as inside(local variable) the function
04. global batiabel can be accessed any function in the same package.
05. local variable only acced by the function in which they declared.
06. variable can be initialised with a value when they are declared,
		or they can be assigned value later.
07. if the variable is not initialised, it'll be assigned zero(0) value for it.
08. the zero value for int is 0(int=0)
	for string ("").
	for boolean (false)
	Short hand variable declaration and assignment.
	:= is used to declare variable,
	when we use this short hand we dont need to specify the type of variable.
	it automatically detect the type based of assigned value.
	i.e.
	age:= 12  // this is int type
	name:= "amit" // this is string type.

	this type of declared variable are not accessable outside the function.
	we can not declare constant(const) type value.
*/
