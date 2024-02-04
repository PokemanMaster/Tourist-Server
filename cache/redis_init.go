package cache

import (
	"github.com/go-redis/redis"
	"os"
	"qiqi-go/middleware/midLogging"
	"strconv"
)

var RDB *redis.Client

// InitRedis 初始化redis
func InitRedis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		midLogging.Info(err)
		panic(err)
	}
	RDB = client
}
