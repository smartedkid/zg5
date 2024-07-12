package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"user_srv/config"
)

var (
	NacosConf  config.NacosConfig
	ServerConf config.ServerConfig
)

var (
	DB  *gorm.DB
	RDB *redis.Client
	Ctx = context.Background()
	ES  *elastic.Client
)
