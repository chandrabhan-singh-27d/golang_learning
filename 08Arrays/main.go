package main

import "fmt"

func main() {
	fmt.Println("Arrays in golang")

	//Array is least used data structure in go
	var fruitList [4]string //there is a necessity to define the number of elements in an array explicitly
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Papaya"

	fmt.Println("Here is our fruitList: ", fruitList)
	//the output of above line is [Apple Banana  Papaya] - we can clearly see there is an extra gap between the second and the last element and that extra space is actually reserved for undefned third element, it means that the space in the memory is allocated even for the undefined element in an array in go

	fmt.Println("The length of the fruit list is: ", len(fruitList))
	//The length of the fruit list is:  4 - this justifies the above statement.

	var vegList = [3]string{"potato", "beans", "mushroom"} //declared and initialized
	fmt.Println("Here is our veg list: ", vegList)
	fmt.Println("The length of the veg list is: ", len(vegList))
}
