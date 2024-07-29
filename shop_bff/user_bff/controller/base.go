package controller

import (
	"0729/shop_bff/user_bff/global"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func ReturnErrJson(err error, c *gin.Context) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  removeTopStruct(errs.Translate(global.Trans)),
		"data": nil,
	})
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
