package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post requests in GO")
	PerformPostJSONRequest()

}

func PerformPostJSONRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Learn Golang",
			"price" : 0,
			"platform" : "Discord"
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
