package sql

import (
	"time"

	"github.com/go-redis/redis"
)

const dashboard = "dashboard"

// Explore ...
func Explore() []string {
	return ExportDB.ZRevRange(dashboard, 0, -1).Val()
}

func addDashboard(RepoName string) {
	ExportDB.ZAdd(dashboard, redis.Z{
		Score:  float64(time.Now().Unix() >> 10),
		Member: RepoName,
	})
}

func delDashboard(RepoName string) {
	ExportDB.ZRem(dashboard, RepoName)
}
