package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to learn slices")

	var fruitList = []string{"Apple", "Papaya", "Guava"} //declare and initialize it
	fmt.Printf("Type of the fruit list is: %T\n", fruitList)
	//Type of the fruit list is: []string - this is a slice
	//Adding elements to the fruitList
	fruitList = append(fruitList, "Mango", "Peaches", "Banana")
	fmt.Println("The fruit list is: ", fruitList)

	//slicing up a slice
	fruitList = append(fruitList[1:]) //start from 1st element with no ending limit
	fmt.Println("The sliced slice is: ", fruitList)
	//The sliced slice is:  [Papaya Guava Mango Peaches Banana]
	fruitList = append(fruitList[1:3])
	fmt.Println("Newly sliced slice is: ", fruitList)
	//Newly sliced slice is:  [Guava Mango] - since we have Guava at the first index of the new slice and the last element is non inclusive therefore we have this new slice of the earlier modified slice
	fruitList = append(fruitList[:1])
	fmt.Println("Slice with no starting value", fruitList)
	//Slice with no starting value [Guava] - started from 0th index of the current slice and returned one element

	highScores := make([]int, 4)
	highScores[0] = 269
	highScores[1] = 854
	highScores[2] = 598
	highScores[3] = 631
	// highScores[4] = 741 - this will throw an error saying index out of range, however, if we append then it will reallocate the memory and each element will be accomodated within that
	highScores = append(highScores, 888, 777, 111) // this is perfectly fine

	fmt.Println(highScores)

	//sorting in go
	sort.Ints(highScores)
	fmt.Println(highScores)                     //sorted slice
	fmt.Println(sort.IntsAreSorted(highScores)) //true

	//Removing a value from a slice based on index

	var courses = []string{"Rust", "Go", "Python", "Ruby", "swift", "javascript"}
	fmt.Println(courses)
	var index int = 4
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) //[Rust Go Python Ruby javascript]
	//what happened here is we started with default index that is zero and included everything till 4 exclusive of 4, then we started with 5(obviously because starting index is inclusive) and ended with whatever the last index is
}
