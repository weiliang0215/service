package initlize

import (
	"github.com/weiliang0215/service/shop_srv/user_srv/global"
	"go.uber.org/zap"
	"net"
)

func InitFreePort() {
	if global.ServerConfig.Port == 0 {
		port, err := GetFreePort()
		if err != nil {
			zap.S().Panic(err)
		}
		global.ServerConfig.Port = port
	}
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	tcp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer tcp.Close()
	return tcp.Addr().(*net.TCPAddr).Port, nil
}
