package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/weiliang0215/service/shop_bff/user_bff/config"
	userPb "github.com/weiliang0215/user_proto/proto"
)

var (
	ServerConfig config.ServerConfig
	RootPath     string
	SrvConnect   userPb.UserClient
	Trans        ut.Translator
)
