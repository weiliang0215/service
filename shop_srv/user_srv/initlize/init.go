package initlize

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/weiliang0215/service/shop_srv/user_srv/global"
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
