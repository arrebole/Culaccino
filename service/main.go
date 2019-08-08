package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/route"
)

func main() {
	server := gin.New()

	// api
	api := server.Group("/api")
	{
		api.GET("/count", route.Count)
		api.GET("/table/:id", route.Table)
		api.GET("/contents/:id", route.Contents)
		api.GET("/login", route.Login)

		api.GET("/token", route.Token)
		api.GET("/admin/table/all", middleware.IsAdmin, route.TableAll)
		api.POST("/admin/add/text", middleware.IsAdmin, route.Add)
		api.PUT("/admin/update/:id", middleware.IsAdmin, route.Update)
		api.DELETE("/admin/delete/:id", middleware.IsAdmin, route.Delete)
	}

	// 静态文件
	server.Use(static.Serve("/", static.LocalFile("./public", true)))
	server.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	// 启动服务器
	server.Run(config.Cofig.Server.ListenPort)
}
