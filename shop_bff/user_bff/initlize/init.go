package initlize

import (
	"github.com/weiliang0215/service/shop_bff/user_bff/global"
	"os"
)

func init() {
	global.RootPath, _ = os.Getwd()

	InitConfig()

	InitFreePort()

	InitLogger()

	InitTranslation("zh")

	InitRegisterValidator()

	InitSrvConnect()
}
