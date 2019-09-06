package sql

import (
	"github.com/go-redis/redis"
	"github.com/arrebole/culaccino/service/config"
)

var instance

func init(){
	instance := redis.NewClient(&redis.Options{
		Addr:     config.DBAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}