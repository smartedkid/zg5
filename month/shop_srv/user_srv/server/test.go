package server

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	"golang.org/x/net/context"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Password  string    `json:"Password"`
	Avatar    string    `json:"avatar"`
	Status    int       `json:"status"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	// 初始化 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	// 存储用户数据到 Redis
	userId := "123"
	userData := map[string]interface{}{
		"ID":        "1",
		"Name":      "Princes Demetris jerde",
		"Mobile":    "1518905132",
		"Password":  "JUIdXTNALCKOICOTDaXWefTORBUKZAKqyVIhyg0bXVpKIndvR",
		"Avatar":    "取细个资",
		"Version":   "v0.0.1",
		"CreatedAt": "2024-07-02T15:33:11.711+08:00",
		"UpdatedAt": "2024-07-11T16:04:25.537+08:00",
	}

	err := rdb.HSet(ctx, "user:"+userId, userData).Err()
	if err != nil {
		log.Fatalf("Failed to set user data: %v", err)
	}

	// 从 Redis 中获取用户数据
	result, err := rdb.HMGet(ctx, "user:"+userId, "ID", "Name", "Mobile", "Password", "Avatar", "Version", "CreatedAt", "UpdatedAt").Result()
	if err != nil {
		log.Fatalf("Failed to get user data: %v", err)
	}

	// 检查获取到的字段数量是否正确
	if len(result) != 8 {
		log.Fatalf("Unexpected result length: %d", len(result))
	}

	// 解析数据并映射到 User 结构体
	user, err := ParseUser(result)
	if err != nil {
		log.Fatalf("Failed to parse user data: %v", err)
	}

	// 打印解析结果
	fmt.Printf("Parsed User: %+v\n", user)
}

func ParseUser(data []interface{}) (User, error) {
	var user User
	var err error

	if user.Id, err = strconv.Atoi(data[0].(string)); err != nil {
		return user, fmt.Errorf("failed to parse ID: %v", err)
	}
	user.Name = data[1].(string)
	user.Mobile = data[2].(string)
	user.Password = data[3].(string)
	user.Avatar = data[4].(string)
	user.Status = data[5].(int)
	user.Version = data[6].(string)

	if user.CreatedAt, err = time.Parse(time.RFC3339, data[6].(string)); err != nil {
		return user, fmt.Errorf("failed to parse CreatedAt: %v", err)
	}

	if user.UpdatedAt, err = time.Parse(time.RFC3339, data[7].(string)); err != nil {
		return user, fmt.Errorf("failed to parse UpdatedAt: %v", err)
	}

	return user, nil
}
