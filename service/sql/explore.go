package sql

import (
	"time"
	"github.com/go-redis/redis"
)

const (
	dashboard = "dashboard"
	maxPerPage = 10
)

// Explore 输出用于首页展示的，repo的名称数组
// 例如 [ "root:dev", "root:test" ]
func (p *SQL) Explore(page int64, perPage int64) []string {
	
	// 限制一次输出的最大值
	perPage = min(maxPerPage, perPage)
	
	// 从ExploreDB数据库的中的， 名称为dashboard的zset结构中读取一定数量的数据
	// dashboard 按照时间顺序排序
	return p.ExploreDB.ZRevRange(dashboard, page*perPage, perPage).Val()
}

// AddExplore 从ExploreDB数据库的中的添加项目,用于首页显示
func (p *SQL) AddExploreItem(RepoName string) {
	p.ExploreDB.ZAdd(dashboard, redis.Z{
		Score:  float64(time.Now().Unix() >> 10),
		Member: RepoName,
	})
}

// 从ExploreDB数据库的中 移除项目
func (p *SQL) DelExploreItem(RepoName string) {
	p.ExploreDB.ZRem(dashboard, RepoName)
}
