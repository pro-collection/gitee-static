package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	file, _ := fileHeader.Open()

	dist := make([]byte, 50000000)                        //开辟存储空间
	n, _ := file.Read(dist)
	sourceString := base64.StdEncoding.EncodeToString(dist[:n])
	c.JSON(http.StatusOK, gin.H{
		"base64": sourceString,
	})
}
