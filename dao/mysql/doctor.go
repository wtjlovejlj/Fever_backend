package mysql

import (
	"Fever_backend/models"
	"go.uber.org/zap"
	"time"
)

//GetMyHosDoctorList  获取本医院所有的医生
func GetMyHosDoctorList(username string, page, size int) (doctors []*models.Doctor, err error) {
	var hos models.HospitalAdmin
	db.Table("hospital_admins").Where("user_name = ?", username).Find(&hos)
	// 查看所有的文章  并分页
	db.Table("doctors").Where("hospital = ?", hos.Hospital).Limit(size).Offset((page - 1) * size).Find(&doctors)
	return
}

//查找所有的医生  并分页
func GetDoctorList(page, size int) (doctors []*models.Doctor, err error) {
	// 查看所有的文章  并分页
	db.Table("doctors").Limit(size).Offset((page - 1) * size).Find(&doctors)
	return
}

//AddDoctor 添加医生
func AddDoctor(up *models.UP, doctor *models.Doctor) (err error) {
	u := db.Table("doctors").Where("username = ?", doctor.Username).Find(doctor)
	if u.RowsAffected > 0 {
		// 用户已存在
		return ErrorUserExit
	}
	// 生成加密密码
	password := encryptPassword([]byte(up.Password))
	db.Table("users").Create(map[string]interface{}{
		"user_name": up.UserName, "password": password, "role": "医生",
	})
	if err != nil {
		zap.L().Error("add  doctor failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	// 把医生插入数据库
	db.Table("doctors").Create(map[string]interface{}{
		"hospital": doctor.Hospital, "id_number": doctor.IDNumber, "phone_number": doctor.PhoneNumber, "realname": doctor.Realname, "username": doctor.Username, "created_at": time.Now(), "district": doctor.District,
	})
	if err != nil {
		zap.L().Error("add  doctor failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

//ChangeDoctorDetailByUserName  修改医生
func ChangeDoctorDetailByUserName(userName string, doctor *models.Doc) (err error) {
	db.Table("doctors").Where("username = ?", userName).Updates(map[string]interface{}{"hospital": doctor.Hospital, "id_number": doctor.IDNumber, "phone_number": doctor.PhoneNumber, "realname": doctor.Realname, "district": doctor.District, "updated_at": time.Now()})
	return err
}

//DeleteDoctorDetailByUserName 删除医生
func DeleteDoctorDetailByUserName(username string) (err error) {
	var doctor *models.Doctor
	u := db.Table("doctors").Where("username = ?", username).Find(&doctor)
	if u.RowsAffected == 0 {
		return ErrorUserNotExit
	}
	db.Table("doctors").Where("username = ?", username).Delete(username)
	db.Table("users").Where("user_name = ?", username).Delete(username)
	return err
}

//UpdateMyMessageByUserName 修改医生管理员信息
func UpdateMyYMessageByUserName(username string, hospitalAdmin *models.Update_my) (err error) {
	db.Table("hospital_admins").Where("user_name = ?", username).Updates(map[string]interface{}{
		"head":  hospitalAdmin.Realname,
		"phone": hospitalAdmin.PhoneNumber,
		"id":    hospitalAdmin.IDNumber,
	})
	return err
}

//UpdateMyMessageByUserName 修改医生个人信息
func UpdateMyDMessageByUserName(username string, hospitalAdmin *models.Update_my) (err error) {
	db.Table("doctors").Where("username = ?", username).Updates(map[string]interface{}{
		"realname":     hospitalAdmin.Realname,
		"phone_number": hospitalAdmin.PhoneNumber,
		"id_number":    hospitalAdmin.IDNumber,
	})
	return err
}
