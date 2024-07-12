package utils

import (
	"fmt"
	"log"
	"net"
	"time"
)

// todo: 获取空闲端口号
func GetFreePort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port
}

func ParseTime(timeStr string) time.Time {
	// 假设你有一个时间字符串，格式为 "2006-01-02T15:04:05Z07:00"
	// 定义时间的布局，它应该与字符串中的格式相匹配
	layout := "2006-01-02T15:04:05Z07:00"

	// 使用 time.Parse 解析时间字符串
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		log.Fatalln("Error parsing time:", err)
	}

	// 打印解析后的时间
	fmt.Println("Parsed time:", parsedTime)
	return parsedTime
}
