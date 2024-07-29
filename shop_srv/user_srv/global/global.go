package global

import (
	"0729/shop_srv/user_srv/config"
	"gorm.io/gorm"
)

var (
	ServerConfig config.ServerConfig
	DB           *gorm.DB
	RootPath     string
	ServerId     string
)
