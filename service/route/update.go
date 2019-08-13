package route

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/arreble/culaccino/service/pareser"
	"github.com/gin-gonic/gin"
)

// Update 更新数据api
func Update(ctx *gin.Context) {
	var (
		id          uint
		postArticle *module.PostArticle
		pares = pareser.New(ctx)
	)

	id, err1 := pares.ParamsID()
	postArticle, err2 := pares.BodyArticle()

	if err1 == nil && err2 == nil {
		sql.New().Update(id, module.ToArticle(postArticle))
		ctx.JSON(200, module.Success())
		return
	}
	ctx.JSON(200, module.Fail())
}
