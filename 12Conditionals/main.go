package main

import "fmt"

func main() {
	fmt.Println("Conditionals in GO")

	loginCount := 23
	var result string

	if loginCount < 10 { //this curly brace cannot be moved to the next line
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	//assigning and checking for the conditions at the same time
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}

	//  if err != nil {
	// 	do something
	// }
}
