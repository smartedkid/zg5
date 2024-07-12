package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user_srv/utils"
)

var err error

func InitReadProfile() {
	path := "./config/config-debug.yaml"
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		//utils.Logger.Info("配置文件读取失败:" + err.Error())
		zap.S().Errorf("配置文件读取失败:%s", err.Error())
		return
	}
	err = viper.Unmarshal(&utils.NacosConf)
	if err != nil {
		//utils.Logger.Info("配置文件解析失败:" + err.Error())
		zap.S().Errorf("配置文件解析失败:%s", err.Error())
		return
	}
	zap.S().Debug("[yaml] 配置文件读取完成")
}
