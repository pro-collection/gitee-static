package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
}
