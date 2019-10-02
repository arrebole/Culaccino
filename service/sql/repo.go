package sql

import (
	"github.com/arrebole/culaccino/service/model"
)

// SetRepo 设置Repo
func (p *SQL) SetRepo(arg *model.Repo) error {
	//作用于首页数据库统计
	p.AddExploreItem(arg.Symbol())
	// 在真正的RepoDB数据库中存放repo
	return p.RepoDB.HMSet(arg.Symbol(), adapter(arg)).Err()
}

// GetRepo ...
func (p *SQL) GetRepo(query string) *model.Repo {
	var result = &model.Repo{}
	reflectMapToStruct(result, p.RepoDB.HGetAll(query).Val())
	return result
}

// GetRepos ...
func (p *SQL) GetRepos(arg ...string) []model.Repo {
	var result = make([]model.Repo, len(arg))
	cmds := getHashPipeline(p.RepoDB, arg...)
	for i := 0; i < len(arg); i++ {
		reflectMapToStruct(&result[i], cmds[i].Val())
	}
	return result
}

// ExistsRepo 检验仓库是否存在
func (p *SQL) ExistsRepo(query string) bool {
	if p.RepoDB.Exists(query).Val() > 0 {
		return false
	}
	return true
}

// DelRepo 删除Repo
// arg -> Repo.Symbol `root:dev`
func (p *SQL) DelRepo(arg string) error {
	// 1、移除repo数据库
	p.RepoDB.Del(arg)
	// 2、移除chapter里所有内容
	// ⚠ TODO

	//4、移除首页
	p.DelExploreItem(arg)

	return nil
}
