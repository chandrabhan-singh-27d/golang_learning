package main

import "fmt"

func main() {
	fmt.Println("Loops in GO")

	//looping over a slice
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	//Conventional syntax and there is no ++i in go
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//Range syntax
	// for j := range days {
	// 	fmt.Println(days[j])
	// }

	//for each syntax -  if you do not want the index replace it with an underscore and remove index as an argument
	for index, day := range days {
		fmt.Printf("index is %v and day is %v\n", index, day)
	}

	rogueValue := 1
	//this is exactly as the while statement
	for rogueValue < 10 {

		if rogueValue == 5 {
			fmt.Println("The value is 5 and We are crossing the permissible limit.")
			rogueValue++
			continue
		}

		if rogueValue == 8 {
			fmt.Println("The value is 8 and we need to break the loop here.")
			break
		}

		if rogueValue == 2 {
			rogueValue++
			goto google
		}

		fmt.Println("Value is :", rogueValue)
		rogueValue++

	google:
		fmt.Println("Jumping at google.com")
	}

}
