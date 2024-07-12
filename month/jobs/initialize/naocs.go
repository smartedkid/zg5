package initialize

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"jobs/utils"
)

func InitReadNacos() {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         utils.NacosConf.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      utils.NacosConf.Host,
			ContextPath: "/nacos",
			Port:        utils.NacosConf.Port,
			Scheme:      "http",
		},
	}
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	//获取配置：GetConfig
	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: utils.NacosConf.DataId,
		Group:  utils.NacosConf.Group})

	err = json.Unmarshal([]byte(content), &utils.ServerConf)
	if err != nil {
		//utils.Logger.Info("[nacos] 配置文件读取失败:" + err.Error())
		zap.S().Errorf("[nacos] 配置文件读取失败:%s", err.Error())
		return
	}
	zap.S().Debug("[nacos] 配置文件读取完成")
}
