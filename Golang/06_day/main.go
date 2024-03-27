package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	//array declararion.
	var fruitList [4]string

	//adding values in list.
	fruitList[0] = "Apple"
	fruitList[1] = "Bananna"
	fruitList[3] = "Mango"

	fmt.Println("fruitList is :", fruitList)
	fmt.Println("fruitList is :", len(fruitList)) // this will give the length of list that is 4
	// but we have inserted only 3 values.
	//instead of giving 3 it give 4 , because we have initialised the size of array ,
	//while declaratin time.
	//this reserve or allocate the 4 memory block, no matter how many value you are inserted.
	// [yes reservation is everywhere]
	// this is fixed size array.[one dimentional]

}
