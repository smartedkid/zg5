package initialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		//utils.Logger.Error("日志初始化失败：" + err.Error())
		zap.S().Errorf("日志初始化失败：%s", err.Error())
		return
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	zap.S().Debug("[logger] 日志初始化完成")
}
