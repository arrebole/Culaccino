package service

import "github.com/go-redis/redis/v7"

var client = conn()

func conn() *redis.Client {
	result := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	testRedis(result)
	return result
}

func testRedis(client *redis.Client) {
	pong, err := client.Ping().Result()
	if pong != "PONG" || err != nil {
		panic("redis conn fail!")
	}
}
