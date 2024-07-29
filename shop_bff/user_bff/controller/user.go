package controller

import (
	"0729/shop_bff/user_bff/forms"
	"0729/shop_bff/user_bff/global"
	"0729/shop_bff/user_bff/server"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	userPb "github.com/weiliang0215/user_proto/proto"
	"net/http"
	"strconv"
	"time"
)

func UserRegister(c *gin.Context) {
	var form forms.UserRegisterForm

	if err := c.ShouldBind(&form); err != nil {
		ReturnErrJson(err, c)
		return
	}
	ctx := context.Background()
	user, err := global.SrvConnect.CreateUser(ctx, &userPb.CreateUserReq{
		Username: form.Username,
		Password: form.Password,
		Mobile:   form.Mobile,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "注册失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": user,
	})
}

func GetUserList(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	if len(page) == 0 || len(limit) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "参数不能为空",
			"data": nil,
		})
		return
	}

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	ctx := context.Background()

	list, err := global.SrvConnect.GetUserList(ctx, &userPb.PageInfoReq{
		Page:  int64(pageInt),
		Limit: int64(limitInt),
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "获取用户列表失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取用户列表成功",
		"data": list,
	})
}

func GetUserInfoByMobile(c *gin.Context) {
	var form forms.UserInfoForm

	if err := c.ShouldBind(&form); err != nil {
		ReturnErrJson(err, c)
		return
	}

	ctx := context.Background()

	userInfo, err := global.SrvConnect.GetUserInfoByMobile(ctx, &userPb.GetUserInfoByMobileReq{Mobile: form.Mobile})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "获取用户信息失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取用户信息成功",
		"data": userInfo,
	})
}

func UserLogin(c *gin.Context) {
	var form forms.UserLoginForm

	if err := c.ShouldBind(&form); err != nil {
		ReturnErrJson(err, c)
		return
	}

	ctx := context.Background()

	userInfo, err := global.SrvConnect.GetUserInfoByMobile(ctx, &userPb.GetUserInfoByMobileReq{Mobile: form.Mobile})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "账号不存在",
			"data": nil,
		})
		return
	}

	check, _ := global.SrvConnect.CheckPassword(ctx, &userPb.CheckPasswordReq{
		Password:           form.Password,
		EncryptionPassword: userInfo.Password,
	})
	if !check.Success {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "密码错误",
			"data": nil,
		})
		return
	}
	marshal, _ := json.Marshal(userInfo)
	token, err := server.SetJwtToken(global.ServerConfig.JwtConfig.Key, time.Now().Unix(), global.ServerConfig.JwtConfig.AccessTokenExpire, string(marshal))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "token生成失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": map[string]interface{}{
			"token":               token,
			"access_token_expire": global.ServerConfig.JwtConfig.AccessTokenExpire,
		},
	})
}

func UserUpdate(c *gin.Context) {
	var form forms.UserUpdateForm

	if err := c.ShouldBind(&form); err != nil {
		ReturnErrJson(err, c)
		return
	}

	ctx := context.Background()
	info, err := global.SrvConnect.UpdateUser(ctx, &userPb.UpdateUserReq{
		Id:    form.Id,
		Email: form.Email,
		Age:   form.Age,
		Sex:   form.Sex,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户信息修改失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户信息修改成功",
		"data": info,
	})

}

func UserDelete(c *gin.Context) {
	var form forms.UserDeleteForm

	if err := c.ShouldBind(&form); err != nil {
		ReturnErrJson(err, c)
		return
	}

	ctx := context.Background()
	user, err := global.SrvConnect.DeleteUser(ctx, &userPb.DeleteUserReq{Mobile: form.Mobile})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户删除失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户删除成功",
		"data": user,
	})
}
