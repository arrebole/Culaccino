package sql

import (
	"github.com/arrebole/culaccino/service/config"
	"github.com/go-redis/redis"
)

// DB 表
var (
	StorageDB    *redis.Client
	MapRepoDB    *redis.Client
	RepoDB       *redis.Client
	MapChapterDB *redis.Client
	ChapterDB    *redis.Client
)

func init() {
	StorageDB = redis.NewClient(&redis.Options{
		Addr:     config.DBAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	MapRepoDB = redis.NewClient(&redis.Options{
		Addr:     config.DBAddr,
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	RepoDB = redis.NewClient(&redis.Options{
		Addr:     config.DBAddr,
		Password: "", // no password set
		DB:       2,  // use default DB
	})
	MapChapterDB = redis.NewClient(&redis.Options{
		Addr:     config.DBAddr,
		Password: "", // no password set
		DB:       3,  // use default DB
	})
	ChapterDB = redis.NewClient(&redis.Options{
		Addr:     config.DBAddr,
		Password: "", // no password set
		DB:       4,  // use default DB
	})

}
