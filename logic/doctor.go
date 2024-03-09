package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"fmt"
	"go.uber.org/zap"
)

//GetDoctorList  获取医生名单列表
func GetDoctorList(username string, page int, size int) (data []*models.Doctor, err error) {
	role, err := mysql.CheckRole(username)
	var er error
	if role == "超级管理员" {
		doctors, err := mysql.GetDoctorList(page, size)
		data = doctors
		er = err
	} else if role == "区县管理员" {
		doctors, err := mysql.MyDistrictDocList(username, page, size)
		data = doctors
		er = err
	} else if role == "院长" {
		doctors, err := mysql.GetMyHosDoctorList(username, page, size)
		data = doctors
		er = err
	}
	return data, er
}

//AddDoctor  添加医生
func AddDoctor(up *models.UP, doctor *models.Doctor) (err error) {
	if err := mysql.AddDoctor(up, doctor); err != nil {
		zap.L().Error("mysql.AddDoctor(doctor) failed", zap.Error(err))
		return err
	}
	return
}

//ChangeDoctorDetail 修改医生信息
func ChangeDoctorDetail(userName string, doctor *models.Doc) error {
	return mysql.ChangeDoctorDetailByUserName(userName, doctor)
}

//DeleteDoctorDetail 删除医生
func DeleteDoctorDetail(username string) error {
	return mysql.DeleteDoctorDetailByUserName(username)
}

//UpdateMyMessage 修改医生管理员的信息
func UpdateMyMessage(username string, doctor *models.Update_my) error {
	role, err := mysql.CheckRole(username)
	fmt.Println(err)
	var a error
	if role == "院长" {
		a = mysql.UpdateMyYMessageByUserName(username, doctor)
	} else if role == "医生" {
		a = mysql.UpdateMyDMessageByUserName(username, doctor)
	}
	return a
}
