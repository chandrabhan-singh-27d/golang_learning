package main

import (
	"fmt"
)

const AccessToken string = "xdnfkjsn564646" //notice that the first letter of the keyword is in capital letter and that indicates that the keyword is a public keyword, in some of the languages the Public keyword needs to be put in order to make the keyword public, here we do it this way.

func main() {
	//String
	/*
		var username string = "Chandrabhan"
		fmt.Println(username)
		fmt.Printf("The variable is of type: %T \n", username)
	*/

	//boolean
	/*
		var isLoggedIn bool = true
		fmt.Println(isLoggedIn)
		fmt.Printf("The variable is of type: %T \n", isLoggedIn)
	*/

	//unsigned integer - can be found from the doc
	/*
		var smallval uint8 = 255
		fmt.Println(smallval)
		fmt.Printf("The variable is of type: %T \n", smallval)
	*/

	//integer
	/*
		var number int = 564789
		fmt.Println(number)
		fmt.Printf("The variable is of type: %T \n", number)
	*/

	//small float, there is a 64 bit floating number as well known as the float64
	/*
		var smallfloat float32 = 255.4546546846513168155
		fmt.Println(smallfloat)
		fmt.Printf("The variable is of type: %T \n", smallfloat)
	*/

	//default values and aliases
	/*
		// var simple_variable int
		var simple_variable string
		fmt.Println(simple_variable)
		fmt.Printf("The variable is of type: %T \n", simple_variable)
	*/

	// this shows the outstanding feature of the golang that it doesn't assign any garbage value as it can be seen in other programming languages

	//Implicit type variable declaration
	/*
		var myName = "Chandra"
		fmt.Println(myName)
		fmt.Printf("The variable is of type: %T \n", myName)
	*/

	//no var type variable declaration - this is allowed inside any method but not outside the method
	/*
		numberOfUser := 789
		fmt.Println(numberOfUser)
		fmt.Printf("The variable is of type: %T \n", numberOfUser)
	*/

	//Accessing a global variable
	fmt.Println(AccessToken)
	fmt.Printf("The variable is of type: %T \n", AccessToken)
}
