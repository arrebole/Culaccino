package sql

import (
	"github.com/arrebole/culaccino/service/module"
)

// Delete 删除
func Delete(arg interface{}) error {
	switch arg.(type) {

	case *module.Repo:
		return deleteRepo(arg.(*module.Repo))

	case *module.Chapter:
		return deleteChapter(arg.(*module.Chapter))

	default:
		panic("arg error")
	}
}

func deleteRepo(arg *module.Repo) error {
	//移除 storage——>repo 表的内容
	MapRepoDB.ZRem(arg.Parents(), arg.Symbol)

	// 2、移除repo数据库
	RepoDB.Del(arg.Symbol)

	// 3、移除chapter里所有内容
	rm, _ := MapChapterDB.ZRange(arg.Symbol, 0, -1).Result()
	for _, v := range rm {
		ChapterDB.Del(v)
	}
	// 4、移除 rpeo-> chapter 的表
	MapChapterDB.Del(arg.Symbol)

	//5、移除首页
	delDashboard(arg.Symbol)

	return nil
}

func deleteChapter(arg *module.Chapter) error {
	MapChapterDB.ZRem(arg.Parents(), arg.Symbol)
	ChapterDB.Del(arg.Symbol)
	return nil
}
