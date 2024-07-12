package from

import "time"

// 用户添加请求参数
type AddUserFrom struct {
	Mobile   string `json:"mobile" validate:"required"`
	Password string `json:"password" validate:"required"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name" validate:"required"`
}

// 用户修改请求参数
type UpdateUserMobile struct {
	UserId int    `json:"user_id" validate:"required,gt>0"`
	Mobile string `json:"mobile" validate:"required"`
}

// 用户删除请求参数
type DeleteUserFrom struct {
	UserId int `json:"user_id" validate:"required,gt>0"`
}

// 用户查询请求参数
type GetUserFrom struct {
	UserId int `json:"user_id" validate:"required,gt>0"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Password  string    `json:"Password"`
	Avatar    string    `json:"avatar"`
	Status    int64     `json:"status"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserOrder struct {
	Name    string  `json:"name"`
	Mobile  string  `json:"mobile"`
	OrderSn string  `json:"order_sn"`
	Status  int64   `json:"status"`
	Total   float64 `json:"total"`
}
