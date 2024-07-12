/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/cobra"
	"strconv"
)

type JSONData struct {
	Action    string `json:"action"`
	Timestamp int    `json:"timestamp"`
	Date      struct {
		Address   string `json:"address"`
		Age       int    `json:"age"`
		Createdat string `json:"createdat"`
		Deletedat string `json:"deletedat"`
		ID        int    `json:"id"`
		Mobile    string `json:"mobile"`
		Sex       int    `json:"sex"`
		Updatedat string `json:"updatedat"`
		Username  string `json:"username"`
	} `json:"date"`
}

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: RedisDel,
}
var (
	Redis *redis.Client
)

func RedisClient() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := Redis.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
func InitRocketmq() {
	topic := "user"
	//clusterName := "DefaultCluster"
	nameSrvAddr := []string{"127.0.0.1:9876"}
	brokerAddr := "127.0.0.1:10911"

	testAdmin, err := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)),
	)

	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr),
	)
	if err != nil {
		fmt.Println("Create topic error:", err.Error())
	}

}
func RedisDel(cmd *cobra.Command, args []string) {

	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("user"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	err := c.Subscribe("user", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		data := JSONData{}
		for _, v := range msgs {
			json.Unmarshal([]byte(v.Body), &data)
		}
		if data.Action == "update" {

			Redis.HDel("user", "user"+strconv.Itoa(data.Date.ID))
		} else {
			fmt.Println("不为删除")
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	select {}
}
func init() {
	rootCmd.AddCommand(jobsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
