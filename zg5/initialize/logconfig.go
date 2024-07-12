package initialize

import "go.uber.org/zap"

func InitZapClient() (*zap.Logger, error) {

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: true,

		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr", "./logs/log.logs"},
		ErrorOutputPaths: []string{"stderr"},
	}
	build, err := config.Build()
	if err != nil {
		return nil, err
	}
	return build, nil
}
