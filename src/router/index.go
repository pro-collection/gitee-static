package router

import (
	"gitee-static/src/controller"
	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controller.Demo)

	return router
}
