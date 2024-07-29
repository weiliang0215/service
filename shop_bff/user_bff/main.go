package main

import (
	"github.com/weiliang0215/service/shop_bff/user_bff/global"
	_ "github.com/weiliang0215/service/shop_bff/user_bff/initlize"
	"github.com/weiliang0215/service/shop_bff/user_bff/router"

	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	initRouter := router.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", global.ServerConfig.IP, global.ServerConfig.Port),
		Handler: initRouter,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

}
