package header

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"zg5/global"
	"zg5/logic"
	"zg5/protoc"
)

type UserServer struct {
	protoc.UnimplementedUserServer
}

func (u *UserServer) GetUserById(ctx context.Context, req *protoc.GetUserByIdReq) (*protoc.GetUserByIdResp, error) {
	resp := protoc.GetUserByIdResp{}
	val := global.RedisClient.HGet("user", "user"+strconv.Itoa(int(req.Id))).Val()
	if val != "" {
		json.Unmarshal([]byte(val), &resp)
		return &resp, nil

	} else {
		user := logic.GetUserById(req)
		copier.Copy(&resp, user)
		marshal, _ := json.Marshal(user)
		global.RedisClient.HSet("user", "user"+strconv.Itoa(int(req.Id)), marshal)
		return &resp, nil
	}
}
func (u *UserServer) AllUser(ctx context.Context, req *protoc.UserAllMessageReq) (*protoc.UserAllMessageResp, error) {
	user := logic.AllUser()
	fmt.Println(user)
	resp := protoc.UserAllMessageResp{Alluser: user}
	return &resp, nil
}
func (u *UserServer) UpdateUser(ctx context.Context, req *protoc.UpdateUserReq) (*protoc.UpdateUserResp, error) {
	err := logic.UpdateUser(req)
	if err != nil {
		return nil, err
	}
	resp := protoc.UpdateUserResp{Success: true}
	return &resp, nil
}
func (u *UserServer) DelUser(ctx context.Context, req *protoc.DelUserReq) (*protoc.DelUserResp, error) {
	err := logic.DelUser(req)
	if err != nil {
		return nil, err
	}
	resp := protoc.DelUserResp{Success: true}
	return &resp, nil
}
func (u *UserServer) AddUser(ctx context.Context, req *protoc.AddUserReq) (*protoc.AddUserResp, error) {
	err := logic.AddUser(req)
	if err != nil {
		return nil, err
	}
	resp := protoc.AddUserResp{Success: true}
	return &resp, nil
}
