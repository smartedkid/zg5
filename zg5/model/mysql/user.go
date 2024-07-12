package mysql

import (
	"time"
)

type User struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	Mobile    string     `json:"mobile"`
	Sex       int        `json:"sex"`
	Age       int        `json:"age"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type Usertables struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Mobile    string    `json:"mobile"`
	Sex       int       `json:"sex"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Userimage string    `json:"userimage"`
}
