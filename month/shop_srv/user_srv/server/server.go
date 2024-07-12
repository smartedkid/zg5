package server

import (
	"errors"
	"fmt"
	"strconv"
	"user_srv/from"
	"user_srv/model/mysql"
	"user_srv/model/redis"
	"user_srv/pkg"
	"user_srv/utils"
)

func GetUserService(data from.GetUserFrom) (*from.User, error) {
	info, err := redis.GetAllHashUser(strconv.FormatInt(int64(data.UserId), 10))
	if err != nil {
		return nil, errors.New("redis查询失败")
	}
	fmt.Println(info)
	if len(info) == 0 {
		fmt.Println(2)
		userInfo, err := mysql.GetUserById(data.UserId)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("mysql查询失败")
		}

		if userInfo.Id == 0 {
			return nil, errors.New("用户数据不存在")
		}

		user := from.User{
			Id:        userInfo.Id,
			Name:      userInfo.Name,
			Mobile:    userInfo.Mobile,
			Password:  userInfo.Password,
			Avatar:    userInfo.Avatar,
			Status:    int64(userInfo.Status),
			Version:   userInfo.Version,
			CreatedAt: userInfo.CreatedAt,
			UpdatedAt: userInfo.UpdatedAt,
		}

		err = redis.AddHashUser(user)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("redis添加失败")
		}

		return &user, nil
	} else {
		fmt.Println(1)
		atoi1, err := strconv.Atoi(info["ID"])
		if err != nil {
			return nil, err
		}
		atoi2, err := strconv.Atoi(info["Status"])
		if err != nil {
			return nil, err
		}
		parseTime1 := utils.ParseTime(info["CreatedAt"])
		parseTime2 := utils.ParseTime(info["UpdatedAt"])
		user := from.User{
			Id:        atoi1,
			Name:      info["Name"],
			Mobile:    info["Mobile"],
			Password:  info["Password"],
			Avatar:    info["Avatar"],
			Status:    int64(atoi2),
			Version:   info["Version"],
			CreatedAt: parseTime1,
			UpdatedAt: parseTime2,
		}

		return &user, nil
	}
}

func AddUserService(data from.AddUserFrom) error {
	user := mysql.Users{
		Name:     data.Name,
		Mobile:   data.Mobile,
		Password: pkg.GetMD5(data.Password),
		Avatar:   data.Avatar,
		Status:   1,
		Version:  "v0.0.1",
	}
	err := mysql.AddUser(&user)
	return err
}

func UpdateUserMobile(data from.UpdateUserMobile) error {
	user := mysql.Users{
		Id:     data.UserId,
		Mobile: data.Mobile,
	}
	err := mysql.UpdateUserMobile(&user)
	return err
}

func DeleteUserService(data from.DeleteUserFrom) error {
	user := mysql.Users{Id: data.UserId}
	err := mysql.DeleteUser(&user)
	return err
}
