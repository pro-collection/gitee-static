package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

type Headers struct {
	AcceptEncoding string `json:"Accept-Encoding"`
}

type User struct {
	Headers Headers `json:"headers"`
}

func main() {

	resp, err := grequests.Get("http://httpbin.org/get", nil)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	user := User{}
	fmt.Println(resp.String())
	_ = resp.JSON(&user)

	fmt.Println(user.Headers.AcceptEncoding)
}
