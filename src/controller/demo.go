package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"log"
)

type Headers struct {
	AcceptEncoding string `json:"Accept-Encoding"`
}

type User struct {
	Headers Headers `json:"headers"`
}

func Demo(c *gin.Context) {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	user := User{}
	_ = resp.JSON(&user)

	c.JSON(200, gin.H{
		"user": user,
	})

}
