package controller

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/logic"
	"Fever_backend/models"
	"Fever_backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//登录
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数 2.校验数据有效性
	var L models.Users
	if err := c.ShouldBindJSON(&L); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	var my models.Update_my
	//用户登录
	if err, a := mysql.Login(&L); err != nil {
		zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	} else {
		my = a
	}

	// 生成Token
	aToken, rToken, _ := jwt.GenToken(L.UserName)
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"username":     L.UserName,
		"role":         L.Role,
		"realname":     my.Realname,
		"phone_number": my.PhoneNumber,
		"id_number":    my.IDNumber,
	})
}

//修改密码
func ChangePasswordHandler(c *gin.Context) {
	//获取参数
	var change_password models.ChangeP
	if err := c.ShouldBindJSON(&change_password); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 获取用户名Name，当前请求的UserName
	userName, err := getCurrentUserName(c)
	if err != nil {
		zap.L().Error("GetCurrentUserName() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	//2.根据name 修改密码
	err = logic.UpdatePasswordDetail(userName, &change_password)
	if err != nil {
		zap.L().Error("mysql.UpdatePasswordDetail() failed", zap.Error(err))
		ResponseError(c, CodeUserExist)
		return
	}
	ResponseSuccess(c, nil)
}
