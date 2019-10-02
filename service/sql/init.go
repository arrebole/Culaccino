package sql

import (
	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/model"
	"github.com/go-redis/redis"
)

var instance *SQL

// SQL 操作客户端
type SQL struct {
	StorageDB *redis.Client
	RepoDB    *redis.Client
	ChapterDB *redis.Client
	ExploreDB *redis.Client
}

func init() {
	instance = &SQL{
		StorageDB: connectDB(0),
		RepoDB:    connectDB(1),
		ChapterDB: connectDB(2),
		ExploreDB: connectDB(3),
	}

	testConnectDB(instance.StorageDB)
	testConnectDB(instance.RepoDB)
	testConnectDB(instance.ChapterDB)
	testConnectDB(instance.ExploreDB)

	defaultDB(instance)
}

// New sql构造函数
func New() Interface {
	return instance
}

func connectDB(index int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.DBAddr,
		Password: "",    // no password set
		DB:       index, // use default DB
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
