package controller

import (
	"Fever_backend/logic"
	"Fever_backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//创建医院管理

func AddHospitalHandler(c *gin.Context) {
	//获取参数
	var up models.UP
	var hospital models.HospitalAdmin
	if err := c.ShouldBindJSON(&hospital); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	up.UserName = hospital.UserName
	up.Password = hospital.Password
	err := logic.AddHospital(&up, &hospital)
	if err != nil {
		zap.L().Error("logic.AddDoctor failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//获取所有医院管理
func DetailHandle(c *gin.Context) {
	userName, err := getCurrentUserName(c)
	if err != nil {
		zap.L().Error("GetCurrentUserName() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	page, size := getPageInfo(c)
	data, err := logic.GetDetailList(userName, page, size)
	if err != nil {
		zap.L().Error("logic.GetBlogList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

//修改医院管理信息
func UpdateDetailHandle(c *gin.Context) {
	Str := c.Query("user_name")
	var hospital models.Hospital
	if err := c.ShouldBindJSON(&hospital); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err := logic.UpdateDetail(Str, &hospital)
	if err != nil {
		zap.L().Error("mysql.ChangeDoctorDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//删除医院管理信息
func DeleteDetailHandle(c *gin.Context) {
	userName := c.Query("user_name")

	err := logic.DeleteDetail(userName)
	if err != nil {
		zap.L().Error("mysql.DeleteDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func ShowDoctorHandle(c *gin.Context) {
	hospital := c.Query("hospital")
	fmt.Println(hospital)

	data, err := logic.ShowDoctor(hospital)
	if err != nil {
		zap.L().Error("mysql.GetUserDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
