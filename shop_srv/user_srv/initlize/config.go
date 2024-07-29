package initlize

import (
	"0729/shop_srv/user_srv/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetSystemConfig() bool {
	// 获取环境变量
	viper.AutomaticEnv()
	getString := viper.GetString("SHOP_MODE")
	fmt.Println(getString)
	if getString == "dev" {
		return false
	} else if getString == "prod" {
		return true
	}
	return false
}

func InitConfig() {
	v := viper.New()

	flag := GetSystemConfig()
	if flag {
		v.SetConfigFile(fmt.Sprintf("%s/config/pro-config.yaml", global.RootPath))
	} else {
		v.SetConfigFile(fmt.Sprintf("%s/config/dev-config.yaml", global.RootPath))
	}
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panic(err)
	}

	go func() {
		v.OnConfigChange(func(e fsnotify.Event) {
			if err := v.Unmarshal(&global.ServerConfig); err != nil {
				zap.S().Panic(err)
			}
			zap.S().Info(global.ServerConfig)
		})
		v.WatchConfig()
	}()

}
