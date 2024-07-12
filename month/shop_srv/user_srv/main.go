package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user_srv/initialize"
	"user_srv/logic"
	"user_srv/proto"
	"user_srv/utils"
)

/*
1.电商商品、库存库表设计
-2.配置viper & nacos，nacos 分组配置
3.git 分支
-4.gorm整合,完成curd、连表查询
-5.日志zap整合
-6.优雅退出&记录日志
-7.注册中心
-8.redis & mysql 读写一致性
-9.cdc mysql to es
-10.es 分词
*/

func main() {
	// 自定义端口号
	port := flag.Int("port", 8080, "端口号")
	if *port == 0 {
		*port = utils.GetFreePort()
	}
	// 初始化zap日志
	initialize.InitLogger()
	// 初始化配置
	initialize.InitReadProfile()
	// 读取nacos配置文件
	initialize.InitReadNacos()
	// 连接mysql
	initialize.InitMysql()
	// 连接redis
	initialize.InitRedis()
	// 连接es
	initialize.InitElasticSearch()

	// 注册grpc服务
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &logic.UserService{})
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		//utils.Logger.Error("监听端口失败:" + err.Error())
		zap.S().Errorf("监听端口:%d失败: %s", *port, err.Error())
		return
	}

	// 注册consul服务
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", utils.ServerConf.Consul.Host, utils.ServerConf.Consul.Port)

	consulClient, err := api.NewClient(config)
	if err != nil {
		//utils.Logger.Error("连接consul失败:" + err.Error())
		zap.S().Errorf("连接consul失败: %s", err.Error())
		return
	}

	// grpc注册服务的健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 每个服务的ID必须不同;这里使用uuid;
	// Name相同ID不同consul会认为是两个实例;
	// Name和ID都相同consul会认为是一个实例会出现覆盖
	serverID := fmt.Sprintf("%s", uuid.NewV4())
	registration := &api.AgentServiceRegistration{
		Address: utils.ServerConf.Host,
		Port:    *port,
		ID:      serverID,
		Name:    utils.ServerConf.Name,
		Tags:    utils.ServerConf.Consul.Tags,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",                                               // 指定运行此检查的频率
			Timeout:                        "5s",                                               // 超时时间
			GRPC:                           fmt.Sprintf("%s:%d", utils.ServerConf.Host, *port), // 健康检查HTTP请求
			DeregisterCriticalServiceAfter: "30s",                                              // 触发注销的时间
		},
	}

	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		//utils.Logger.Error("注册服务失败:" + err.Error())
		zap.S().Errorf("注册服务失败: %s", err.Error())
		return
	}
	zap.S().Debugf("[grpc] 服务启动成功;PORT:%d", *port)

	go func() {
		// 服务连接
		if err = server.Serve(listen); err != nil && err != http.ErrServerClosed {
			//utils.Logger.Error("监听失败:" + err.Error())
			zap.S().Errorf("监听失败: %s", err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println(time.Now().Format(time.DateTime))

	err = consulClient.Agent().ServiceDeregister(serverID)
	if err != nil {
		//utils.Logger.Error("注销失败:" + err.Error())
		zap.S().Errorf("注销失败: %s", err.Error())
		return
	}
	//utils.Logger.Info("[grpc] 服务注销成功")
	zap.S().Debug("[grpc] 服务注销成功")
}
