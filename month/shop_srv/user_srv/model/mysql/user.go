package mysql

import (
	"gorm.io/gorm"
	"user_srv/utils"
)

// todo: 用户表
type Users struct {
	Id       int    `gorm:"primary"`
	Name     string `gorm:"column:name;type:varchar(255);comment:账号"`
	Mobile   string `gorm:"column:mobile;type:char(11);comment:手机号"`
	Password string `gorm:"column:password;type:varchar(255);comment:密码"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);comment:头像"`
	Status   int    `gorm:"column:status;type:tinyint;comment:用户状态"`
	Version  string `gorm:"column:version;type:varchar(12);comment:版本号"`
	gorm.Model
}

// 用户添加
func AddUser(user *Users) (err error) {
	err = utils.DB.Create(&user).Error
	return
}

// 修改用户手机号
func UpdateUserMobile(user *Users) (err error) {
	err = utils.DB.Model(&user).Where("id = ?", user.Id).Update("mobile", user.Mobile).Error
	return
}

// 用户删除（软删）
func DeleteUser(user *Users) (err error) {
	err = utils.DB.Where("id = ?", user.Id).Delete(&user).Error
	return err
}

func GetUserById(userId int) (user *Users, err error) {
	err = utils.DB.Where("id = ?", userId).Find(&user).Limit(1).Error
	return
}
