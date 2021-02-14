package main

import (
	"gitee-static/src/router"
)

func main() {
	app := router.SetUp()
	_ = app.Run(":8082") // 监听并在 0.0.0.0:8080 上启动服务
}
