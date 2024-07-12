package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zg5/global"
)

func InitMysqlClient() {
	host := global.SerceConfig.Mysql.Host
	port := global.SerceConfig.Mysql.Port
	user := global.SerceConfig.Mysql.User
	password := global.SerceConfig.Mysql.Password
	db := global.SerceConfig.Mysql.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, db)
	var err error
	global.MysqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql 链接失败")
	} else {
		fmt.Println("mysql 连接成功")
	}
}
