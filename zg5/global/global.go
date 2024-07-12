package global

import (
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
	"zg5/config"
)

var (
	NacosConfig config.NacosConfig
	SerceConfig config.ServerConfig
	MysqlDb     *gorm.DB
	RedisClient *redis.Client
)
