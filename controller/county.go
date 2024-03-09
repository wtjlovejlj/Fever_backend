package controller

import (
	"Fever_backend/logic"
	"Fever_backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//ChangeCountyHandler 修改区县
func ChangeCountyHandler(c *gin.Context) {
	oldcounty := c.Query("district")
	var con models.County
	if err := c.ShouldBindJSON(&con); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err := logic.ChangeCountyDetail(oldcounty, &con)
	if err != nil {
		zap.L().Error("mysql.ChangeCountyDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//DeleteCountyHandler 删除区县
func DeleteCountyHandler(c *gin.Context) {
	county := c.Query("district")
	err := logic.DeleteCountyDetail(county)
	if err != nil {
		zap.L().Error("mysql.DeleteCountyDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//FindCountyHandler  查找区县的所有医院
func FindCountyHandler(c *gin.Context) {
	county := c.Query("district")
	//获取数据
	data, err := logic.FindCountyList(county)
	if err != nil {
		zap.L().Error("Logic.FindCountyList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

//DistrictListHandler  区县展示
func DistrictListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	//获取数据
	data, err := logic.GetDistrictList(page, size)
	if err != nil {
		zap.L().Error("Logic.GetDistrictList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}
