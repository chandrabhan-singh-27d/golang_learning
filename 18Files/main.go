package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in GO")
	content := "This needs to go in a file - google.com"

	file, err := os.Create("./myGoFile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length of the file is: ", length)
	defer file.Close() //recommended as this needs to be done at the end of the operation
	readFile("./myGoFile.txt")
	//Text data inside the file is
	// [84 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105 110 32 97 32 102 105 108 101 32 45 32 103 111 111 103 108 101 46 99 111 109]
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is \n", databyte)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
1. Whenever you read a file it will be read in a byte format and not the string.

*/
