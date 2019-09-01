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
		api.GET("/session/login", controllers.SessionLogin)
		api.GET("/session/exist", controllers.SessionExist)

		api.POST("/static/upload", middleware.Auth, controllers.StaticUpload)

		api.POST("/archive/create", middleware.Auth, controllers.ArchiveCreate)
		api.PUT("/archive/update/:id", middleware.Auth, controllers.ArchiveUpdate)
		api.DELETE("/archive/delete/:id", middleware.Auth, controllers.ArchiveDelete)
		api.GET("/archive/details/:id", controllers.ArchivesDetails)
		api.GET("/archive/all", middleware.Auth, controllers.ArchiveAll)
		api.GET("/archive/dashboard/:page", controllers.ArchiveDashboard)

	}

	// 静态文件
	router.Use(static.Serve("/", static.LocalFile("./wwwroot/", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	return router
}
