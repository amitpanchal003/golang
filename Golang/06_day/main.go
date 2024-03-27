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
	fmt.Println("this is all about Slices")
	var vegiList = []string{"amit", "sumit", "samit"} // slice having three values(items)
	//whatever datatype we mention, it become that data type
	// here we mentioned []string, so it become string datatype of datastructure of that list
	//in slices, no need to define the size of slices
	// slices are like D-array
	// while declaring slices we have to initializing it same time.

	fmt.Printf("type of list is %T\n", vegiList)
	fmt.Println("Before adding any other values: ", vegiList)
	//lets add more items in list.
	vegiList = append(vegiList, "mit", "mYth")
	// using append method adding iteams in list
	// append method require listName and values[items]
	fmt.Println("After adding new values: ", vegiList)
	fmt.Println("########################")
	highScore := make([]int, 4)

	highScore[0] = 564
	highScore[1] = 102
	highScore[2] = 201
	highScore[3] = 674
	// now try to add more values in list
	//highScore[4] = 787
	//fmt.Println(highScore)// it"ll give panc runtime error : index out of bound.

	// try to add values using append method
	highScore = append(highScore, 999, 888, 777)
	fmt.Println(highScore) // after addding print the list.
	//thsi will work smoothly.
	// bcoz, when we use append method it will re-allocate the memory and store the values
	// this is special part in golang slices
	fmt.Print("size of highScore is : ", len(highScore)) // size is 7
	// it automatically increase the size of slices according to the list values.
}
