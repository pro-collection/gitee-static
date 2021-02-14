package controller

import (
	"encoding/base64"
	"fmt"
	"gitee-static/src/config"
	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"net/http"
	"time"
)

func Upload(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	file, _ := fileHeader.Open()

	dist := make([]byte, 50000000) //开辟存储空间
	n, _ := file.Read(dist)
	sourceString := base64.StdEncoding.EncodeToString(dist[:n])
	//c.JSON(http.StatusOK, gin.H{
	//	"base64": sourceString,
	//})

	fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano()/1e6, fileHeader.Filename)

	// 请求
	requestUrl := fmt.Sprintf(
		"%s%s",
		"https://gitee.com/api/v5/repos/yanleweb/static/contents/hd_client/",
		fileName,
	)
	//	https://gitee.com/api/v5/repos/yanleweb/static/contents/hd_client/demo-2.jpg

	res, err := grequests.Post(requestUrl, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Length": "multipart/form-data; boundary=<calculated when request is sent>",
		},
		Data: map[string]string{
			"access_token": config.Get().AccessToken,
			"message":      fileName,
			"content":      sourceString,
		},
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"base64":       sourceString,
		"download_url": gjson.Get(res.String(), "content.download_url").String(),
	})

}
