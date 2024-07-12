package initialize

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"zg5/global"
)

func InitRedisClient() {
	host := global.SerceConfig.Redis.Host
	port := global.SerceConfig.Redis.Port
	password := global.SerceConfig.Redis.Password
	db := global.SerceConfig.Redis.Db
	addr := fmt.Sprintf("%s:%s", host, port)
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	pong, err := global.RedisClient.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
