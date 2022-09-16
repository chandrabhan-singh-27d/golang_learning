package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("World")
	deferFunc()
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}
}

/*
1. The defer keyword defer the statement at put it at the end of the function same as we defer the script in html.
2. When we have multiple defer statements the LIFO model will be applied.
*/
