package logic

import (
	"context"
	"errors"
	"fmt"
	"time"
	"user_srv/from"
	"user_srv/model/mysql"
	"user_srv/model/redis"
	"user_srv/pkg"
	"user_srv/proto"
	"user_srv/server"
)

// 用户信息查询
func (u *UserService) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	userId := in.UserId

	data := from.GetUserFrom{
		UserId: int(userId),
	}

	service, err := server.GetUserService(data)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		Id:      int64(service.Id),
		Name:    service.Name,
		Mobile:  service.Mobile,
		Avatar:  service.Avatar,
		Status:  service.Status,
		Version: service.Version,
	}, nil
}

func (u *UserService) GetUserList(ctx context.Context, in *proto.GetUserListRequest) (*proto.GetUserListResponse, error) {
	//user, err := redis.GetAllHashUser("1")
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(user)
	// 入参接受
	name := in.Name
	mobile := in.Mobile
	passwd := in.Password
	avatar := in.Avatar
	// 参数校验
	data := from.AddUserFrom{
		Mobile:   mobile,
		Password: passwd,
		Avatar:   avatar,
		Name:     name,
	}
	fmt.Println(data)
	UserData := from.User{
		Id:        3,
		Name:      name,
		Mobile:    mobile,
		Password:  passwd,
		Avatar:    avatar,
		Status:    1,
		Version:   "v0.0.1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err := redis.AddHashUser(UserData)
	if err != nil {
		return nil, errors.New("添加失败" + err.Error())
	}

	return &proto.GetUserListResponse{
		Success: true,
	}, nil
}

// 用户添加
func (u *UserService) AddUser(ctx context.Context, in *proto.AddUserRequest) (*proto.AddUserResponse, error) {
	// 入参接受
	name := in.Name
	mobile := in.Mobile
	passwd := in.Password
	avatar := in.Avatar
	// 参数校验
	data := from.AddUserFrom{
		Mobile:   mobile,
		Password: passwd,
		Avatar:   avatar,
		Name:     name,
	}

	err := pkg.ValidateFroms(data)
	if err != nil {
		return nil, errors.New("参数校验失败" + err.Error())
	}

	err = server.AddUserService(data)
	if err != nil {
		return nil, errors.New("用户添加失败失败" + err.Error())
	}

	// 参数返回值
	return &proto.AddUserResponse{
		Success: true,
	}, nil
}

// 修改用户手机号
func (u *UserService) UpdateUser(ctx context.Context, in *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	userId := in.UserId
	mobile := in.Mobile

	data := from.UpdateUserMobile{
		UserId: int(userId),
		Mobile: mobile,
	}

	err := pkg.ValidateFroms(data)
	if err != nil {
		return &proto.UpdateUserResponse{
			Success: false,
		}, errors.New("参数校验失败" + err.Error())
	}

	err = server.UpdateUserMobile(data)
	if err != nil {
		return &proto.UpdateUserResponse{
			Success: false,
		}, errors.New("手机号修改失败" + err.Error())
	}

	return &proto.UpdateUserResponse{
		Success: true,
	}, nil
}

// 删除用户
func (u *UserService) DeleteUser(ctx context.Context, in *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	userId := in.UserId

	data := from.DeleteUserFrom{
		UserId: int(userId),
	}

	err := server.DeleteUserService(data)
	if err != nil {
		return nil, errors.New("用户信息删除失败失败" + err.Error())
	}

	return &proto.DeleteUserResponse{
		Success: true,
	}, nil
}

// 查询用户是否存在订单
func (u *UserService) GetUserOrder(ctx context.Context, in *proto.GetUserOrderRequest) (*proto.GetUserOrderResponse, error) {
	userId := in.UserId
	order, err := mysql.GetOrder(int(userId))
	if err != nil {
		return nil, err
	}
	fmt.Println(order)
	return &proto.GetUserOrderResponse{
		UserId:      userId,
		UserName:    order.Name,
		UserMobile:  order.Mobile,
		OrderSn:     order.OrderSn,
		OrderStatus: order.Status,
		OrderAmount: float32(order.Total),
	}, nil

}
