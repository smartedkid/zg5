package initialize

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"user_srv/utils"
)

func InitElasticSearch() {
	esInfo := utils.ServerConf.Es
	addr := fmt.Sprintf("http://%s:%s", esInfo.Host, esInfo.Port)
	utils.ES, err = elastic.NewClient(
		elastic.SetURL(addr),
		elastic.SetSniff(false),
	)
	if err != nil {
		//utils.Logger.Error("[elasticsearch] 初始化失败：" + err.Error())
		zap.S().Errorf("[elasticsearch] 初始化失败：%s", err.Error())
		return
	}
	zap.S().Debugf("[elasticsearch] 连接成功 >>>>>>>>>>>>>>> PORT:%s", esInfo.Port)
}
