package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weiliang0215/service/shop_bff/user_bff/controller"
	"github.com/weiliang0215/service/shop_bff/user_bff/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("user")
	{
		// 用户注册
		v1.POST("", controller.UserRegister)
		// 用户列表
		v1.GET("", controller.GetUserList)
		// 用户信息
		v1.GET("/mobile", middleware.MiddleToken(), controller.GetUserInfoByMobile)
		// 用户登录
		v1.POST("/login", controller.UserLogin)
		// 用户修改
		v1.PUT("", controller.UserUpdate)
		// 用户删除
		v1.DELETE("", controller.UserDelete)
	}

	return router

}
