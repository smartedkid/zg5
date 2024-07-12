package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerFile() *zap.Logger {
	c := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "M",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "./log/log.log"},
		ErrorOutputPaths: []string{"stderr", "./log/error.log"},
	}
	build, err := c.Build()
	if err != nil {
		return nil
	}
	return build
}
