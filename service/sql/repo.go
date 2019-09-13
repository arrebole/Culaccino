package sql

import (
	"errors"
	"strings"
	"github.com/arrebole/culaccino/service/model"
	"github.com/go-redis/redis"
)

// SetRepo 设置Repo
func (p *SQL) SetRepo(arg *model.Repo) error {
	// 从MapRepoDB中添加 storage -> repo 的索引
	p.SetRepoMap(arg.Symbol())
	
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

// SetRepoMap ...
// arg -> repo.Symbol() `root:dev`
func (p *SQL) SetRepoMap(arg string) error {
	split := strings.Split(arg, ":")
	if len(split) != 2 {
		return errors.New("arg need xxx:xxx")
	}

	return p.MapRepoDB.ZAdd(split[0], redis.Z{
		Score:  0,
		Member: arg,
	}).Err()
}

// DelRepoMap ...
// arg: repo.Symbol() `root:dev`
func (p *SQL) DelRepoMap(arg string) error {
	split := strings.Split(arg, ":")
	if len(split) != 2{
		return errors.New("arg need xxx:xxx")
	}
	return p.MapRepoDB.ZRem(split[0], arg).Err()
}

// GetRepoMap 获取storage对应的repos的key数组
func (p *SQL) GetRepoMap(arg string) []string {
	return p.MapRepoDB.ZRange(arg, 0, -1).Val()
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
	//移除 storage——>repo 表的内容
	p.DelRepoMap(arg)

	// 2、移除repo数据库
	p.RepoDB.Del(arg)

	// 3、移除chapter里所有内容
	rmItems := p.GetChapterMap(arg)
	for _, v := range rmItems {
		p.DelChapter(v)
	}

	//4、移除首页
	p.DelExploreItem(arg)

	return nil
}