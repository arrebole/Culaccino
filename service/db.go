package service

import "github.com/go-redis/redis/v7"

var client = conn()

func conn() *redis.Client {
	result := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// test conn
	if err := testConn(result); err != nil {
		panic(err.Error())
	}
	return result
}

func testConn(client *redis.Client) error {
	return client.Ping().Err()
}
