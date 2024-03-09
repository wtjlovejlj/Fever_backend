package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"Fever_backend/pkg/snowflake"
	"go.uber.org/zap"
)

//创建信息
func CreateMessage(fever *models.FeverPeople) (err error) {
	// 生成信息ID
	feverID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}
	fever.FeverID = feverID

	// 创建信息
	if err := mysql.CreatePost(fever); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}

	return

}

//GetFareList  展示列表
func GetFareList(page, size int) (data []*models.FeverPeople, err error) {
	Fare, err := mysql.GetFareList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetDoctorList(page,size) failed", zap.Error(err))
		return
	}
	data = Fare
	return
}

//修改信息
func UpdateMessage(id int64, fever *models.FeverPeople) error {
	return mysql.UpdateMessage(id, fever)
}

//删除信息
func DeleteMsg(id int64) error {
	return mysql.DeleteMsgByID(id)
}
