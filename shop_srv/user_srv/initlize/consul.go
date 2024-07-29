package initlize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/weiliang0215/service/shop_srv/user_srv/global"

	"go.uber.org/zap"
)

func InitConsul() {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.ConsulHost, global.ServerConfig.ConsulConfig.ConsulPort)

	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Panic(err)
	}
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", global.ServerConfig.IP, global.ServerConfig.Port),
		Timeout:                        "5s",
		Interval:                       "2s",
		DeregisterCriticalServiceAfter: "1s",
	}

	consul := new(api.AgentServiceRegistration)
	consul.Name = global.ServerConfig.ServerName
	consul.ID = global.ServerId
	consul.Port = global.ServerConfig.Port
	consul.Address = global.ServerConfig.IP
	consul.Tags = global.ServerConfig.ConsulConfig.ConsulTags
	consul.Check = check

	err = client.Agent().ServiceRegister(consul)
	if err != nil {
		zap.S().Panic(err)
	}
	fmt.Println("ok")

}
