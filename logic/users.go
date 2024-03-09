package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
)

func UpdatePasswordDetail(userName string, change_password *models.ChangeP) error {
	return mysql.UpdatePasswordDetailByName(userName, change_password)
}
