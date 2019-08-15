package route

import (
	"github.com/arrebole/culaccino/service/controllers"
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// New 创建路由
func New() *gin.Engine {

	router := gin.New()
	// api
	api := router.Group("/api")
	{
		api.GET("/table/:page", controllers.Table)
		api.GET("/archives/:id", controllers.Archives)

		api.GET("/login", controllers.Login)

		api.GET("/token", controllers.Token)
		api.GET("/admin/table", middleware.Auth, controllers.TableAll)
		api.POST("/admin/add", middleware.Auth, controllers.Add)
		api.PUT("/admin/update/:id", middleware.Auth, controllers.Update)
		api.DELETE("/admin/delete/:id", middleware.Auth, controllers.Delete)

		api.GET("/admin/database/download", middleware.Auth, controllers.DBDownLoad)
		api.POST("/admin/database/upload", middleware.Auth, controllers.DBUpload)
	}

	// 静态文件
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	return router
}
