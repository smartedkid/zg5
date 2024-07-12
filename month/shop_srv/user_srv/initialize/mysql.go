package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"user_srv/utils"
)

func InitMysql() {
	c := utils.ServerConf.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
	fmt.Println(dsn)
	utils.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		//utils.Logger.Error("[mysql] 初始化失败：" + err.Error())
		zap.S().Errorf("[mysql] 初始化失败：%s", err.Error())
		return
	}

	zap.S().Debugf("[mysql] 连接成功 >>>>>>>>>>>>>>> PORT:%s", c.Port)
}
