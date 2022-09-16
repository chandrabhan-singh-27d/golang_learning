package main

import "fmt"

func main() {
	fmt.Println("Understanding structs in GO")
	//There is no inheritance in GO; No super or parent/child

	chandra := User{"Chandrabhan", "chandra@go.dev", true, 27}
	fmt.Println(chandra)
	//printing with details
	fmt.Printf("chandra details are: %+v\n", chandra)
	//getting specifics
	fmt.Printf("Name of the person is %v and their email is %v\n", chandra.Name, chandra.Email)
}

// capitalized syntax as things need to be exported
type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
}
