package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"user_srv/from"
	"user_srv/utils"
)

const UserPrifx = "USER_DATA"
const User = "UserId_"

//------------------------------------------string-----------------------------------------------------------------------

func AddUserRedis(userId string, data from.AddUserFrom) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = utils.RDB.Set(utils.Ctx, UserPrifx+User+userId, marshal, 86420*time.Second).Err()
	return err
}

func GetUserInfo(userId string) (string, error) {
	result, err := utils.RDB.Get(utils.Ctx, UserPrifx).Result()
	return result, err
}

func GetUser(userId string) (string, error) {
	result, err := utils.RDB.Get(utils.Ctx, UserPrifx+User+userId).Result()
	return result, err
}

func CheckUser(userId string) (int64, error) {
	result, err := utils.RDB.Exists(utils.Ctx, User+userId).Result()
	return result, err
}

func DelUser(userId string) error {
	err := utils.RDB.Del(utils.Ctx, User+userId).Err()
	return err
}

//------------------------------------------hash-----------------------------------------------------------------------

func AddHashUser(data from.User) error {
	//var userData []map[string]interface{}
	Data := map[string]interface{}{
		"ID":        data.Id,
		"Name":      data.Name,
		"Mobile":    data.Mobile,
		"Password":  data.Password,
		"Avatar":    data.Avatar,
		"Version":   data.Version,
		"Status":    data.Status,
		"CreatedAt": data.CreatedAt,
		"UpdatedAt": data.UpdatedAt,
	}
	//userData = append(userData, Data)
	hashKey := fmt.Sprintf("%s::%s", UserPrifx, User+strconv.Itoa(data.Id))
	//for _, datum := range userData {
	err := utils.RDB.HMSet(utils.Ctx, hashKey, Data).Err()
	return err
	//}
	//return nil
}

//func AddHashUser(data from.User) error {
//	marshal, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//	hashKey := fmt.Sprintf("%s:%s", User, strconv.Itoa(data.Id))
//	userInfo := map[string]string{
//		hashKey: string(marshal),
//	}
//	err = utils.RDB.HMSet(utils.Ctx, UserPrifx, userInfo).Err()
//	return err
//}

func GetHashUser(userId string) ([]interface{}, error) {
	result, err := utils.RDB.HMGet(utils.Ctx, UserPrifx, UserPrifx+userId).Result()
	return result, err
}

//func GetHashUser(userId string) ([]interface{}, error) {
//	result, err := utils.RDB.HMGet(utils.Ctx, UserPrifx, User+userId, "Id", "Name", "Mobile", "Password", "Avatar", "Status", "Version", "CreatedAt", "UpdatedAt").Result()
//	return result, err
//}

func GetAllHashUser(userId string) (map[string]string, error) {
	hashKey := fmt.Sprintf("%s::%s", UserPrifx, User+userId)
	//hashKey2 := UserPrifx + ":"
	result, err := utils.RDB.HGetAll(utils.Ctx, hashKey).Result()
	return result, err
}
