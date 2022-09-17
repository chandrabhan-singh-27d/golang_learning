package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Requests in GO")
	response, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Response is of type %T\n", response) // Response is of type *http.Response - this is a ponter so we are not getting the copy of the response but the response itself

	defer response.Body.Close() //It is callers responsibility to close the connection

	databyte, err := io.ReadAll(response.Body)
	checkNilError(err)
	content := string(databyte)
	fmt.Println(content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
