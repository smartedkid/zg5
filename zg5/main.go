package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"zg5/global"
	"zg5/header"
	"zg5/initialize"
	"zg5/protoc"
	"zg5/utils"
)

func main() {
	port := flag.Int("port", 8888, "http to userServer")
	flag.Parse()
	logger, err := initialize.InitZapClient()
	if err != nil {
		return
	}
	logger.Sync()
	initialize.InitNacosClient()
	initialize.InitMysqlClient()
	initialize.InitRedisClient()
	if *port == 0 {
		*port = utils.GenFreePort()
	}
	server := grpc.NewServer()
	//proto.RegisterTestServer(server, &service.Service{})
	protoc.RegisterUserServer(server, &header.UserServer{})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听端口:%d失败: %s", *port, err.Error())
	}

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", global.SerceConfig.Consul.Host, global.SerceConfig.Consul.Port)

	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("连接consul失败: %s", err.Error())
	}

	// grpc注册服务的健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 每个服务的ID必须不同;这里使用uuid;
	// Name相同ID不同consul会认为是两个实例;
	// Name和ID都相同consul会认为是一个实例会出现覆盖
	registration := &api.AgentServiceRegistration{
		Address: global.SerceConfig.Host,
		Port:    *port,
		ID:      fmt.Sprintf("%s-%s-%d", global.SerceConfig.Host, global.SerceConfig.Name, *port),
		Name:    global.SerceConfig.Name,
		Tags:    global.SerceConfig.Tags,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",                                                        // 指定运行此检查的频率
			Timeout:                        "5s",                                                        // 超时时间
			GRPC:                           fmt.Sprintf("%s:%d", global.SerceConfig.Consul.Host, *port), // 健康检查HTTP请求
			DeregisterCriticalServiceAfter: "30s",                                                       // 触发注销的时间
		},
	}
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("注册服务失败: %s", err.Error())
	}

	fmt.Printf("服务启动成功;PORT:%d\n", *port)
	go func() {
		err = server.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("成功退出" + strconv.Itoa(*port))
	log.Println("成功退出")

}
