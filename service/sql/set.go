package sql

import (
	"errors"

	"github.com/arrebole/culaccino/service/module"
	"github.com/go-redis/redis"
)

// Set 创建
func Set(arg interface{}) error {
	switch arg.(type) {
	case *module.Storage:
		return setStorage(arg.(*module.Storage))
	case *module.Repo:
		return setRepo(arg.(*module.Repo))
	case *module.Chapter:
		return setChapter(arg.(*module.Chapter))
	default:
		return errors.New("")
	}
}

func setStorage(arg *module.Storage) error {
	return hmset(StorageDB, arg.Name, adapter(arg))
}

func setRepo(arg *module.Repo) error {
	addmap(MapRepoDB, arg.Parents, arg.Name)
	return hmset(RepoDB, arg.Name, adapter(arg))
}

func setChapter(arg *module.Chapter) error {
	addmap(MapChapterDB, arg.Parents, arg.Name)
	return hmset(ChapterDB, arg.Name, adapter(arg))
}

func hmset(db *redis.Client, key string, data map[string]interface{}) error {
	return db.HMSet(key, data).Err()
}

// 增加映射数据库
func addmap(db *redis.Client, src string, dist string) error {
	return db.ZAdd(src, redis.Z{
		Score:  0,
		Member: dist,
	}).Err()
}
