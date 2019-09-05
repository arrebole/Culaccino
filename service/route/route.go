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

		api.GET("/search/dashboard", controllers.SearchDashboard)

		api.POST("/new", controllers.RepoNew)
		api.GET("/repo/:domain", controllers.RepoDomain)
		api.GET("/repo/:domain/:symbol", controllers.RepoDetails)
		api.PUT("/repo/:domain/:symbol", controllers.RepoCommit)
		api.DELETE("/repo/:domain/:symbol", controllers.RepoDelete)

	}

	// 静态文件
	router.Use(static.Serve("/", static.LocalFile("./wwwroot/", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	return router
}
