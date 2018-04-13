package lib

import (
	"strconv"
	"wechatHimsAPI/config"

	"github.com/go-redis/redis"
)

const KeyHead = "wechat_hims_api:"

var RedisClient *redis.Client

func InitRedisClient() {
	redisDB, err := strconv.Atoi(config.C["redis_db"])
	if err != nil {
		panic("config: redis_db is empty!")
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.C["redis_host"] + ":" + config.C["redis_port"],
		Password: config.C["redis_pwd"], // no password set
		DB:       redisDB,               // use default DB
	})
}
