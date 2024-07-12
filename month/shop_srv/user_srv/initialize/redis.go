package initialize

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"user_srv/utils"
)

func InitRedis() {
	addr := utils.ServerConf.Redis.Host + ":" + utils.ServerConf.Redis.Port
	utils.RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err = utils.RDB.Ping(utils.Ctx).Result()
	if err != nil {
		//utils.Logger.Error("[redis] 初始化失败：" + err.Error())
		zap.S().Errorf("[redis] 初始化失败：%s", err.Error())
		return
	}
	zap.S().Debugf("[redis] 连接成功 >>>>>>>>>>>>>>> PORT:%s", utils.ServerConf.Redis.Port)
}
