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

		api.POST("/new", controllers.Create)

		api.GET("/dashboard", controllers.GetDashboard)
		api.GET("/search", controllers.Search)
		api.POST("/upload", controllers.StaticUpload)
		api.PUT("/commit", controllers.Commit)
		api.DELETE("/delete", controllers.Delete)

		api.GET("/session", controllers.Session)

	}

	// 静态文件
	router.Use(static.Serve("/", static.LocalFile("./wwwroot/", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	return router
}
