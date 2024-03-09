package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"go.uber.org/zap"
)

//创建医院管理
func AddHospital(up *models.UP, hospital *models.HospitalAdmin) (err error) {
	if err := mysql.AddHospital(up, hospital); err != nil {
		zap.L().Error("mysql.AddDoctor(doctor) failed", zap.Error(err))
		return err
	}
	return
}

//修改
func UpdateDetail(user string, hospitalAdmin *models.Hospital) error {
	return mysql.UpdateDetailByCard(user, hospitalAdmin)
}

func ShowDoctor(hospital string) ([]models.Doctor, error) {
	return mysql.ShowDoctorByHospital(hospital)
}

//分页
func GetDetailList(username string, page, size int) (data []*models.HospitalAdmin, err error) {
	role, err := mysql.CheckRole(username)
	var er error
	if role == "超级管理员" {
		hos, err := mysql.GetAllList(page, size)
		data = hos
		er = err
	} else if role == "区县管理员" {
		hos, err := mysql.MyDistrictHosList(username, page, size)
		data = hos
		er = err
	}
	return data, er
}

//删除
func DeleteDetail(UserName string) error {
	return mysql.DeleteDetailByCard(UserName)
}
