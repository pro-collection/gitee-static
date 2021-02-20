# gitee-static

白嫖 gitee

gitee 在国内访问非常快，可以考虑把一些静态资源上传到到 gitee 上面， 然后直接访问， 岂不是可以当做静态资源服务器来使用？

###  具体实现
参考 [gitee open api](https://gitee.com/api/v5/swagger#/postV5ReposOwnerRepoContentsPath)

```go
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

	dist := make([]byte, fileHeader.Size) //开辟存储空间
	n, _ := file.Read(dist)
	sourceString := base64.StdEncoding.EncodeToString(dist[:n])

	fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano()/1e6, fileHeader.Filename)
	urlPath := fmt.Sprintf(
		"https://gitee.com/api/v5/repos/%s/%s/contents/%s/",
		config.Get().Owner,
		config.Get().Repository,
		config.Get().Path,
	)

	// 请求
	requestUrl := fmt.Sprintf(
		"%s%s",
		urlPath,
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
```

需要在 config 目录下面创建 config.yaml 文件， 配置自己的仓库配置信息.                      
配置的具体内容可以参考上面 gitee openApi 就行
```yaml
access_token: "*****"
owner: "***"
repository: "****"
path: "****"
```


启动项目之后访问 `127.0.0.1:8082` 上传文件即可使用


