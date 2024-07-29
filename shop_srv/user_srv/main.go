package main

import (
	"0729/shop_srv/user_srv/global"
	"0729/shop_srv/user_srv/handler"
	_ "0729/shop_srv/user_srv/initlize"
	"fmt"
	"github.com/hashicorp/consul/api"

	userPb "github.com/weiliang0215/user_proto/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	zap.S().Info(global.ServerConfig)

	g := grpc.NewServer()

	s := handler.UserService{}

	grpc_health_v1.RegisterHealthServer(g, health.NewServer())

	userPb.RegisterUserServer(g, &s)

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", global.ServerConfig.IP, global.ServerConfig.Port))
	if err != nil {
		zap.S().Panic(err)
	}

	go func() {
		err = g.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	config := api.DefaultConfig()

	config.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.ConsulHost, global.ServerConfig.ConsulConfig.ConsulPort)

	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Panic(err)
	}

	err = client.Agent().ServiceDeregister(global.ServerId)
	if err != nil {
		zap.S().Info("服务注销失败")
		return
	}

	zap.S().Info("服务注销成功")
}
