package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbewi432jfg"

func main() {
	fmt.Println("Handling URLs in GO")
	fmt.Println(myurl)

	//parsing url - extracting some meaningful info
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) //url.Values
	fmt.Println(qparams)

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=chnara",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
