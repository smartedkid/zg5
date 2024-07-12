/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"jobs/model/redis"
	"jobs/utils"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mysql redis 数据一致")
		MQConsumeRedis()
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func MQConsumeRedis() {
	Consume()
}

// 主题模式消费消息
func Consume() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("group1"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)

	err = c.Subscribe("users", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		var data utils.Data
		for _, msg := range msgs {
			fmt.Printf("%s,%s", msg.Body, "------------------------------------")
			err = json.Unmarshal(msg.Body, &data)
			if err != nil {
				fmt.Println("json格式转换失败")
			}
		}
		fmt.Println(data)
		switch {
		case data.Action == "update":
			err = redis.DelHashUser(strconv.Itoa(data.Date.ID))
			if err != nil {
				fmt.Println("缓存删除失败", err.Error())
			}
		default:
			fmt.Println("无需缓存删除操作")
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	select {}
}
