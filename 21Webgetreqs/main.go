package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Handling get requests in GO")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	//handling error on upcoming response
	if err != nil {
		panic(err)
	}

	//close the call
	defer response.Body.Close()

	//check for the status code
	fmt.Println("The status code is: ", response.StatusCode)
	//check for the content lenght
	fmt.Println("The content length is: ", response.ContentLength)
	content, _ := io.ReadAll(response.Body)
	/*
		//Read the content from the body
		//Print the content in the string format
		fmt.Println(string(content))
	*/
	//Reading and throwing the data in the string format can also be done using the strings package as shown below

	var responseString strings.Builder
	// byteCount, _ := responseString.Write(content)
	// fmt.Println("ByteCount is: ", byteCount)
	responseString.Write(content)
	fmt.Println(responseString.String())
}
