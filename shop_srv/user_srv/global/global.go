package global

import (
	"github.com/weiliang0215/service/shop_srv/user_srv/config"
	"gorm.io/gorm"
)

var (
	ServerConfig config.ServerConfig
	DB           *gorm.DB
	RootPath     string
	ServerId     string
)
