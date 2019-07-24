package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/arrebole/culaccino/src/config"
	"github.com/arrebole/culaccino/src/route"
)

func main() {
	server := gin.New()

	// api
	api := server.Group("/api")
	{
		api.GET("/table/:page", route.Table)
		api.GET("/contents/:id", route.Contents)
		api.GET("/login", route.Login)

		api.POST("/add/text", route.Add)
		api.PUT("/update/:id", route.Update)
		api.DELETE("/delete/:id", route.Delete)
	}

	// 静态文件
	server.NoRoute(static.Serve("/", static.LocalFile("./public", true)))

	// 启动服务器
	server.Run(config.Cofig.Server.ListenPort)
}
