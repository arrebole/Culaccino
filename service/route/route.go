package route

import (
	"github.com/arrebole/culaccino/service/controllers"
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

		api.POST("/static/upload", controllers.StaticUpload)

		api.GET("/archive/details/:id", controllers.ArchivesDetails)
		api.POST("/archive/create", controllers.ArchiveCreate)
		api.PUT("/archive/update/:id", controllers.ArchiveUpdate)
		api.DELETE("/archive/delete/:id", controllers.ArchiveDelete)
		api.GET("/archive/owner", controllers.ArchiveOwner)
		api.GET("/archive/dashboard/:page", controllers.ArchiveDashboard)

	}

	// 静态文件
	router.Use(static.Serve("/", static.LocalFile("./wwwroot/", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	return router
}
