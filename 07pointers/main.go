package main

import "fmt"

func main() {
	fmt.Println("Welcome to the tut on pointers")

	//declaring a pointer without referencing it will give a nil value
	/*
		var ptr *int
		fmt.Println("The value of the pointer is", ptr)
	*/

	//referencing a pointer
	myNumber := 89
	var ptr = &myNumber
	fmt.Println("The value of the pointer is", ptr)
	fmt.Println("The value of the refernced variable is", *ptr)

	//mutating the value of referenced variable using pointers
	*ptr = *ptr + 3
	fmt.Println("The new value of the reference variable is: ", myNumber)
}

//pointers gives the access to the memory location meaning the memory address is being passed giving the full surety that the actual value is being passed and not a copy of the variable. sometimes, when we pass the value of a variable, the copy of the value is passed and it may lead to problems and in order to avoid such situations we make use of the pointers
