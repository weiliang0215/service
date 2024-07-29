package initlize

import (
	"github.com/weiliang0215/service/shop_srv/user_srv/global"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	flag := GetSystemConfig()

	if flag {
		cfg := zap.NewProductionConfig()

		cfg.OutputPaths = []string{
			global.ServerConfig.LogConfig.Path,
		}
		return cfg.Build()
	} else {
		cfg, err := zap.NewDevelopment()
		return cfg, err
	}
}

func InitLogger() {
	logger, err := NewLogger()

	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}
