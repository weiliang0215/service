package initlize

import (
	"0729/shop_srv/user_srv/global"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
)

func init() {
	global.RootPath, _ = os.Getwd()
	fmt.Println(global.RootPath)
	InitConfig()

	global.ServerId = fmt.Sprintf("%s", uuid.NewV4())

	InitFreePort()

	InitLogger()

	InitMysql()

	InitConsul()

}
