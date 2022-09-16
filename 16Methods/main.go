package main

import "fmt"

func main() {
	fmt.Println("Methods in GO")

	chandra := User{"Chandrabhan", "chandra@go.dev", true, 27}
	chandra.GetStatus()
	chandra.ResetEmail()

	//check if it changes the actual property
	fmt.Printf("The original email of the user is %v\n", chandra.Email)
}

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
}

func (u User) GetStatus() {
	fmt.Println("Is user verified: ", u.IsVerified)
}

func (u User) ResetEmail() {
	u.Email = "chandra2@go.dev"
	fmt.Println("The new email of the user is: ", u.Email)
}

/*
1. Functions when brought inside the structs are known as Methods as there are no classes in GO.
2. Syntax
	func(reciver_name Type) method_name(parameter_list)(return_type){
		do something
	}
3. The structs are passed as a copy and not the actual values thus we have pointers designed.
4. Whenever we want to pass the actual reference of the structs we should pass them as the reference variables and not the actual values
*/
