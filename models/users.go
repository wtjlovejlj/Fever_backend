package models

import (
	"encoding/json"
	"errors"
)

//权限结构
type CasbinModel struct {
	RoleName string `json:"role" db:"v0"`
	Path     string `json:"path" db:"v1"`
	Method   string `json:"method" db:"v2"`
}

func (L *CasbinModel) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		RoleName string `json:"role" db:"v0"`
		Path     string `json:"path" db:"v1"`
		Method   string `json:"method" db:"v2"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else {
		L.RoleName = required.RoleName
		L.Path = required.Path
		L.Method = required.Method
	}
	return
}

type UP struct {
	UserName string `json:"username" db:"username"` // 用户名
	Password string `json:"password" db:"password"` // 密码
}

type Users struct {
	UserName string `json:"username" db:"username"` // 用户名
	Role     string `json:"role" db:"role"`         // 用户角色
	Password string `json:"password" db:"password"` // 密码
}

func (L *Users) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		L.UserName = required.UserName
		L.Password = required.Password
	}
	return
}

type ChangeP struct {
	Old_password     string `json:"oldpassword" db:"oldpassword"`           // 旧密码
	New_password     string `json:"newpassword" db:"newpassword"`           // 新密码
	Confirm_password string `json:"confirm_password" db:"confirm_password"` // 确认密码
}

func (Cp *ChangeP) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Old_password     string `json:"oldpassword" db:"oldpassword"`           // 旧密码
		New_password     string `json:"newpassword" db:"newpassword"`           // 新密码
		Confirm_password string `json:"confirm_password" db:"confirm_password"` // 确认密码
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Old_password) == 0 {
		err = errors.New("缺少必填字段oldpassword")
	} else if len(required.New_password) == 0 {
		err = errors.New("缺少必填字段newpassword")
	} else if len(required.Confirm_password) == 0 {
		err = errors.New("缺少必填字段confirm_password")
	} else if required.Confirm_password != required.New_password {
		err = errors.New("两次输入的密码不一致")
	} else {
		Cp.Old_password = required.Old_password
		Cp.New_password = required.New_password
		Cp.Confirm_password = required.Confirm_password
	}
	return
}
