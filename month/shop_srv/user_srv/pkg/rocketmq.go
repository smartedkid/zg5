package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"user_srv/from"
	"user_srv/model/mysql"
)

// 创建主题
func InitRocketmq(str string) {
	topic := str
	nameSrvAddr := []string{"127.0.0.1:9876"}
	brokerAddr := "127.0.0.1:10911"

	testAdmin, err := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)),
	)

	//create topic
	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr),
	)
	if err != nil {
		fmt.Println("Create topic error:", err.Error())
	}
}

// 主题模式消费消息
func Consume() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)

	err = c.Subscribe("users", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe callback: %v \n", msgs)

		var message from.AddUserFrom

		for _, msg := range msgs {
			err = json.Unmarshal(msg.Body, &message)
			if err != nil {
				return 0, err
			}
		}

		data := mysql.Users{
			Name:     message.Name,
			Mobile:   message.Mobile,
			Password: message.Password,
			Avatar:   message.Avatar,
			Status:   1,
			Version:  "v.0.0.1",
		}

		err = mysql.AddUser(&data)
		if err != nil {
			fmt.Println(err.Error())
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

// 主题模式生产消息
func Produce(data from.AddUserFrom) {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)

	if err != nil {
		fmt.Println("init producer error: " + err.Error())
		os.Exit(0)
	}

	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	res, err := p.SendSync(context.Background(), primitive.NewMessage(
		"users",
		jsonData),
	)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}

	err = p.Shutdown()
	if err != nil {
		fmt.Printf("topic生产消息失败: %s", err.Error())
	}
}
