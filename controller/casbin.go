package controller

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AddCasbin(c *gin.Context) {
	var X models.CasbinModel
	if err := c.ShouldBindJSON(&X); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	isok := mysql.AdCasbin(X)
	if isok {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "保存成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "保存失败",
		})
	}

}
