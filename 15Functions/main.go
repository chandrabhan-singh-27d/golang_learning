package main

import "fmt"

func main() {
	fmt.Println("Functions in GO")
	greetInSpanish()
	greetInEnglish()
	result := adder(5, 9)
	fmt.Println("The result is: ", result)
	proResult, myMessage := proAdder(3, 4, 5, 23, 56, 4, 7)
	fmt.Println("The pro result is: ", proResult)
	fmt.Println("My message is: ", myMessage)
}

func greetInSpanish() {
	fmt.Println("Hola in espanol!!")
}

func greetInEnglish() {
	fmt.Println("Hello in english!")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi Pro Adder Function"
}

/*
Note:
1. We are not allowed to define functions inside main() it expects statements only.
2.The order of the function definition does not matter, it's the calling of the function that always matters.
3. Anonymous functions do exists and we can create IIFEs as well in GO.
4. return type aka signature is expected if we are returning something out of a function.
5. VarArgs also work in GO as described in the proAdder()
6. It can have multiple return type as well.
*/
