package router

import (
	"gitee-static/src/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	api := router.Group("/api")
	{
		api.GET("/ping", controller.Demo)
		api.POST("/upload", controller.Upload)
	}

	template := router.Group("/")
	{
		template.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", nil)
		})
	}

	return router
}
