package sql

import (
	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/model"
	"github.com/go-redis/redis"
)

var instance *SQL

type SQL struct {
	StorageDB    *redis.Client
	MapRepoDB    *redis.Client
	RepoDB       *redis.Client
	MapChapterDB *redis.Client
	ChapterDB    *redis.Client
	ExploreDB    *redis.Client
}

func init() {
	instance = &SQL{
		StorageDB:    connectDB(0),
		MapRepoDB:    connectDB(1),
		RepoDB:       connectDB(2),
		MapChapterDB: connectDB(3),
		ChapterDB:    connectDB(4),
		ExploreDB:    connectDB(5),
	}

	testConnectDB(instance.StorageDB)
	testConnectDB(instance.MapRepoDB)
	testConnectDB(instance.RepoDB)
	testConnectDB(instance.MapChapterDB)
	testConnectDB(instance.ChapterDB)
	testConnectDB(instance.ExploreDB)

	defaultDB(instance)
}

func New() Interface {
	return instance
}

func connectDB(index int)*redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.DBAddr,
		Password: "",     // no password set
		DB:       index,  // use default DB
	})
}

func testConnectDB(p *redis.Client) {
	if p.Ping().Val() != "pong" {
		panic("数据库连接失败")
	}
}

func defaultDB(db Interface) {
	if !db.ExistsStorage("root") {
		db.SetStorage(&model.Storage{
			Name:     "root",
			Password: config.PassWord,
		})
	}
}
