package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Handling form data requests in GO")
}

func PerformFormDataPostRequest() {
	const myUrl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstName", "Chandra")
	data.Add("lastName", "Singh")
	data.Add("email", "chandra@go.dev")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
