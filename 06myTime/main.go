package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of go lang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//change format - as per the doc we need to use this format only
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//create time
	createdDate := time.Date(2020, time.December, 27, 00, 00, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

//build can be made for any platform(os) - the command for that can be seen from go env
//to build an executable file for a different os than your current env we can do it as GOOS="osname" go build where osname=windows or linux or darwin(for mac)
