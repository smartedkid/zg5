package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"zg5/global"
)

func InitNacosClient() {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	v.ReadInConfig()
	nacosConfig := global.NacosConfig
	v.Unmarshal(&nacosConfig)
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.Host,
			Port:   uint64(nacosConfig.Port),
		}}
	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group})
	json.Unmarshal([]byte(content), &global.SerceConfig)
	fmt.Println(global.SerceConfig)
}
