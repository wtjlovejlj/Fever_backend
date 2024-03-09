package mysql

import (
	"Fever_backend/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const secret = "fever.com"

//密码加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

//登录
func Login(user *models.Users) (err error, my models.Update_my) {
	originPassword := user.Password
	u := db.Table("users").Where("user_name = ?", user.UserName).Find(&user)
	if u.RowsAffected == 0 {
		// 用户不存在
		return ErrorUserNotExit, my
	}
	var use models.Users
	db.Table("users").Select("password").Where("user_name = ?", user.UserName).Find(&use)
	password := encryptPassword([]byte(originPassword))
	if use.Password != password {
		return ErrorPasswordWrong, my
	}
	var m1 models.Update_my
	var m2 models.Update_my2
	role, err := CheckRole(user.UserName)
	fmt.Println(err)
	fmt.Println(user.UserName)
	if role == "院长" {
		db.Table("hospital_admins").Where("user_name = ?", user.UserName).Find(&m2)
		my.Realname = m2.Head
		my.IDNumber = m2.ID
		my.PhoneNumber = m2.Phone
	} else if role == "医生" {
		db.Table("doctors").Where("username = ?", user.UserName).Find(&m1)
		my = m1
	}
	return err, my
}

//修改密码
func UpdatePasswordDetailByName(userName string, change_password *models.ChangeP) (err error) {
	originPassword := change_password.Old_password
	//var cp models.ChangeP
	var user models.Users
	db.Table("users").Select("password").Where("user_name = ?", userName).Find(&user)
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return ErrorPasswordWrong
	}
	change_password.New_password = encryptPassword([]byte(change_password.New_password))

	db.Table("users").Where("user_name = ?", userName).Updates(map[string]interface{}{"password": change_password.New_password})

	return err
}

func CheckRole(usernme string) (role string, err error) {
	var use *models.Users
	db.Table("users").Where("user_name = ?", usernme).Find(&use)
	return use.Role, err
}
