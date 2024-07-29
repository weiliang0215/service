package global

import (
	"0729/shop_bff/user_bff/config"
	ut "github.com/go-playground/universal-translator"
	userPb "github.com/weiliang0215/user_proto/proto"
)

var (
	ServerConfig config.ServerConfig
	RootPath     string
	SrvConnect   userPb.UserClient
	Trans        ut.Translator
)
