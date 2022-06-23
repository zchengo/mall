package initialize

import (
	"github.com/go-redis/redis/v8"
	"imall/global"
)

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	global.Rdb = rdb
}
