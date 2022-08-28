package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please rate our pizza")

	//comma ok syntax - used in case of try/catch, not exactly same but somewhere close.
	input, _ := reader.ReadString('\n') //using _ as a variable is possible but we should try to avoid it
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of the input is %T \n", input)

}
