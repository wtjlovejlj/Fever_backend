package mysql

import (
	"Fever_backend/models"
	"go.uber.org/zap"
)

//建表
func AddHospital(up *models.UP, hospital *models.HospitalAdmin) (err error) {
	u := db.Table("hospital_admins").Where("user_name = ?", hospital.UserName).Find(hospital)
	if u.RowsAffected > 0 {
		// 用户已存在
		return ErrorUserExit
	}
	// 生成加密密码
	password := encryptPassword([]byte(up.Password))
	db.Table("users").Create(map[string]interface{}{
		"user_name": up.UserName, "password": password, "role": "院长",
	})
	if err != nil {
		zap.L().Error("add  hospitalAdmin failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	// 把医生插入数据库
	db.Table("hospital_admins").Create(map[string]interface{}{
		"credit":    hospital.Credit,
		"phone":     hospital.Phone,
		"id":        hospital.ID,
		"district":  hospital.District,
		"hospital":  hospital.Hospital,
		"user_name": hospital.UserName,
		"head":      hospital.Head,
		"address":   hospital.Address,
	})
	if err != nil {
		zap.L().Error("add  doctor failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

//展示所有的医院管理信息，包括分页
func GetAllList(page, size int) (posts []*models.HospitalAdmin, err error) {
	db.Table("hospital_admins").Limit(size).Offset((page - 1) * size).Find(&posts)
	return
}

//修改医院管理信息
func UpdateDetailByCard(user_name string, hospitalAdmin *models.Hospital) (err error) {
	db.Table("hospital_admins").Where("user_name = ?", user_name).Updates(map[string]interface{}{
		"district": hospitalAdmin.District,
		"hospital": hospitalAdmin.Hospital,
		"credit":   hospitalAdmin.Credit,
		"address":  hospitalAdmin.Address,
		"head":     hospitalAdmin.Head,
		"phone":    hospitalAdmin.Phone,
		"id":       hospitalAdmin.ID,
	})
	return err
}

//删除医院管理信息
func DeleteDetailByCard(UserName string) (err error) {
	db.Table("hospital_admins").Where("user_name = ?", UserName).Delete(UserName)
	db.Table("users").Where("user_name = ?", UserName).Delete(UserName)
	return err
}

func ShowDoctorByHospital(Hospital string) (doctor []models.Doctor, err error) {
	db.Table("doctors").Where("hospital = ?", Hospital).Find(&doctor)
	return doctor, err
}
