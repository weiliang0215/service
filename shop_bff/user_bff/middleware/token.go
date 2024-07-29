package middleware

import (
	"0729/shop_bff/user_bff/global"
	"0729/shop_bff/user_bff/model"
	"0729/shop_bff/user_bff/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "token不可以为空",
				"data": nil,
			})
		}
		payload, err := server.ParseToken(global.ServerConfig.JwtConfig.Key, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "token时效性过期",
				"data": nil,
			})
		}
		user := model.User{}
		json.Unmarshal([]byte(payload), &user)

		c.Set("user_id", user.ID)

		c.Next()
	}
}
