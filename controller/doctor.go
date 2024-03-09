package controller

import (
	"Fever_backend/logic"
	"Fever_backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//DoctorListHandler 获取医生名单的处理函数
func DoctorListHandler(c *gin.Context) {
	userName, err := getCurrentUserName(c)
	if err != nil {
		zap.L().Error("GetCurrentUserName() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	page, size := getPageInfo(c)
	//获取数据
	data, err := logic.GetDoctorList(userName, page, size)
	if err != nil {
		zap.L().Error("Logic.GetDoctorList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

//AddDoctorHandler  添加医生
func AddDoctorHandler(c *gin.Context) {
	//获取参数
	var up models.UP
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	up.UserName = doctor.Username
	up.Password = doctor.Password
	err := logic.AddDoctor(&up, &doctor)
	if err != nil {
		zap.L().Error("logic.AddDoctor failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//ChangeDoctorHandler 修改医生信息
func ChangeDoctorHandler(c *gin.Context) {
	userName := c.Query("username")
	var doctor models.Doc
	if err := c.ShouldBindJSON(&doctor); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err := logic.ChangeDoctorDetail(userName, &doctor)
	if err != nil {
		zap.L().Error("mysql.ChangeDoctorDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//DeleteDoctorHandler 删除医生
func DeleteDoctorHandler(c *gin.Context) {
	userName := c.PostForm("username")
	//2.根据id 获取博客详情
	err := logic.DeleteDoctorDetail(userName)
	if err != nil {
		zap.L().Error("mysql.DeleteDoctorDetail() failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}
	ResponseSuccess(c, nil)
}

//UpdateMyMessage 修改当前医生管理员的信息
func UpdateMyMessage(c *gin.Context) {
	// 获取作者Name，当前请求的UserName
	userName, err := getCurrentUserName(c)
	if err != nil {
		zap.L().Error("GetCurrentUserName() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	var my models.Update_my
	if err := c.ShouldBindJSON(&my); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err = logic.UpdateMyMessage(userName, &my)
	if err != nil {
		zap.L().Error("mysql.UpdateMyMessage() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
