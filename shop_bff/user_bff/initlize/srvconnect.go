package initlize

import (
	"0729/shop_bff/user_bff/global"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	userPb "github.com/weiliang0215/user_proto/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConnect() {
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulConfig.ConsulHost, global.ServerConfig.ConsulConfig.ConsulPort, global.ServerConfig.ConsulConfig.UserSrvName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Panic(err)
	}

	global.SrvConnect = userPb.NewUserClient(conn)
}
