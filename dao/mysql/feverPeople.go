package mysql

import (
	"Fever_backend/models"
	"go.uber.org/zap"
	"time"
)

// CreatePost 创建发热人员信息
func CreatePost(fever *models.FeverPeople) (err error) {
	u := db.Where("id_card_id = ?", fever.IDCardID).Find(fever)

	if u.RowsAffected > 0 {
		// 信息已存在
		return ErrorMsgExit
	}

	// 把信息插入数据库
	db.Table("fever_peoples").Create(map[string]interface{}{
		"fever_id":          fever.FeverID,
		"address":           fever.Address,
		"age":               fever.Age,
		"clinical_effect":   fever.ClinicalEffect,
		"color_code":        fever.ColorCode,
		"disposal_method":   fever.DisposalMethod,
		"doctor":            fever.Doctor,
		"id_card_id":        fever.IDCardID,
		"mobile_phone":      fever.MobilePhone,
		"name":              fever.Name,
		"nucleic_acid_test": fever.NucleicAcidTest,
		"primaryi_dagnosis": fever.PrimaryiDagnosis,
		"program_results":   fever.ProgramResults,
		"sex":               fever.Sex,
		"date":              time.Now(),
	})

	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

//得到信息列表
func GetFeverDetailByFeverID(NameStr string) (fever []models.FeverPeople, err error) {
	db.Table("FeverPeople").Select("fever_id", "address",
		"age",
		"clinical_effect",
		"color_code",
		"disposal_method",
		"doctor",
		"id_card_id",
		"mobile_phone",
		"name",
		"nucleic_acid_test",
		"primaryi_dagnosis",
		"program_results",
		"sex",
		"creat_at",
	).Where("Name = ?", NameStr).Find(&fever)
	return fever, err
}

// 查看所有的信息  并分页
func GetFareList(page, size int) (posts []*models.FeverPeople, err error) {
	db.Table("fever_peoples").Limit(size).Offset((page - 1) * size).Find(&posts)
	return
}

//修改信息
func UpdateMessage(id int64, fever *models.FeverPeople) (err error) {
	db.Table("fever_peoples").Where("fever_id = ?", id).Updates(map[string]interface{}{
		"address":           fever.Address,
		"age":               fever.Age,
		"clinical_effect":   fever.ClinicalEffect,
		"color_code":        fever.ColorCode,
		"disposal_method":   fever.DisposalMethod,
		"doctor":            fever.Doctor,
		"id_card_id":        fever.IDCardID,
		"mobile_phone":      fever.MobilePhone,
		"name":              fever.Name,
		"nucleic_acid_test": fever.NucleicAcidTest,
		"primaryi_dagnosis": fever.PrimaryiDagnosis,
		"program_results":   fever.ProgramResults,
		"sex":               fever.Sex,
		"date":              time.Now(),
	})
	return err
}

//删除信息
func DeleteMsgByID(id int64) (err error) {
	db.Table("fever_peoples").Where("fever_id = ?", id).Delete(id)
	return err
}
