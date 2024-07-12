package logic

import (
	"fmt"
	"github.com/jinzhu/copier"
	"time"
	"zg5/global"
	"zg5/model/mysql"
	"zg5/protoc"
)

func GetUserById(req *protoc.GetUserByIdReq) mysql.User {
	user := mysql.User{}
	global.MysqlDb.Table("user").Where("id=?", req.Id).First(&user)
	return user
}
func AllUser() []*protoc.AllUser {
	var user []mysql.Usertables
	global.MysqlDb.Table("user").Select("user. *, usertables.userimage").Joins("  join usertables on user.id=usertables.userid ").Scan(&user)
	var allUser []*protoc.AllUser
	for _, v := range user {
		p := protoc.AllUser{}
		fmt.Println(v)
		copier.Copy(&p, v)
		allUser = append(allUser, &p)
	}
	return allUser
}
func UpdateUser(req *protoc.UpdateUserReq) error {

	err := global.MysqlDb.Exec("update user set  sex = ? where id=?", req.Sex, req.Id).Error
	if err != nil {
		return err
	}
	return nil
}
func DelUser(req *protoc.DelUserReq) error {
	err := global.MysqlDb.Exec("update user set  deleted_at  = current_date where id=?", req.Id).Error
	if err != nil {
		return err
	}
	return nil
}
func AddUser(req *protoc.AddUserReq) error {

	user := mysql.User{
		Username:  req.Username,
		Mobile:    req.Mobile,
		Sex:       int(req.Sex),
		Age:       int(req.Age),
		Address:   req.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := global.MysqlDb.Table("user").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
