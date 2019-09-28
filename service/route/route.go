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

		api.POST("/new/repo", controllers.NewRepo)
		api.POST("/new/chapter", controllers.NewChapter)

		api.GET("/dashboard", controllers.GetDashboard)
		api.GET("/search", controllers.Search)
		api.POST("/upload", controllers.StaticUpload)

		api.PUT("/commit/repo", controllers.CommitRepo)
		api.PUT("/commit/chapter", controllers.CommitChapter)

		api.DELETE("/delete/repo", controllers.DelRepo)

		api.GET("/session/exists", controllers.SessionExists)
		api.GET("/session/login", controllers.SessionLogin)

	}

	// 静态文件
	router.Use(static.Serve("/", static.LocalFile("./wwwroot/", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	return router
}
