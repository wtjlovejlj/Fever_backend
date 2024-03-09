package controller

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/logic"
	"Fever_backend/models"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//  AddMessage 创建信息列表
func AddMessage(c *gin.Context) {

	//
	var fever models.FeverPeople

	if err := c.ShouldBindJSON(&fever); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	err := logic.CreateMessage(&fever)

	if errors.Is(err, mysql.ErrorMsgExit) {
		ResponseError(c, CodeUserExist)
		return
	}

	if err != nil {
		zap.L().Error("logic.CreateBlog failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

//发热人员列表
func FareListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	//获取数据
	data, err := logic.GetFareList(page, size)
	if err != nil {
		zap.L().Error("Logic.GetFareList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

//DeleteMessage获取信息
func DeleteMessage(c *gin.Context) {
	//1.获取信息ID
	idStr := c.Param("feverId") //获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	//2.根据id 获取信息详情
	err = logic.DeleteMsg(id)
	if err != nil {
		zap.L().Error("mysql.DeleteMsg() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// UpdateMessage修改信息
func UpdateMessage(c *gin.Context) {
	//1.获取信息ID
	idStr := c.Param("feverId") //获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	var fever models.FeverPeople
	if err := c.ShouldBindJSON(&fever); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	//2.根据id 获取信息详情
	err = logic.UpdateMessage(id, &fever)
	if err != nil {
		zap.L().Error("mysql.UpdateMessage() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
