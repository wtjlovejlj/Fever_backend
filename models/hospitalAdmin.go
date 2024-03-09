package models

import (
	"encoding/json"
	"errors"
)

type HospitalAdmin struct {
	Card   int64  `json:"card" db:"card" gorm:"primaryKey"`
	Credit string `json:"credit" db:"credit"`
	Phone  string `json:"phone" db:"phone"`
	ID     string `json:"ID" db:"ID"`

	District string `json:"district" db:"district"`
	Hospital string `json:"hospital" db:"hospital"`
	Head     string `json:"head" db:"head"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	Address  string `json:"address" db:"address"`
}

type Hospital struct {
	Credit string `json:"credit" db:"credit"`
	Phone  string `json:"phone" db:"phone"`
	ID     string `json:"ID" db:"ID"`

	District string `json:"district" db:"district"`
	Hospital string `json:"hospital" db:"hospital"`
	Head     string `json:"head" db:"head"`
	UserName string `json:"user_name" db:"user_name"`
	Address  string `json:"address" db:"address"`
}

type Update_my struct {
	IDNumber    string `json:"id_number" db:"id_number"`       // 身份证号
	PhoneNumber string `json:"phone_number" db:"phone_number"` // 电话号码
	Realname    string `json:"realname" db:"realname"`         // 真实姓名
}
type Update_my2 struct {
	Head  string `json:"head" db:"head"`
	Phone string `json:"phone" db:"phone"`
	ID    string `json:"ID" db:"ID"`
}

func (d *Update_my) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		IDNumber    string `json:"id_number" db:"id_number"`       // 身份证号
		PhoneNumber string `json:"phone_number" db:"phone_number"` // 电话号码
		Realname    string `json:"realname" db:"realname"`         // 真实姓名
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.IDNumber) != 18 {
		err = errors.New("身份证号格式错误")
	} else if len(required.PhoneNumber) != 11 {
		err = errors.New("手机号格式错误")
	} else if len(required.Realname) == 0 {
		err = errors.New("真是姓名不能为空")
	} else {
		d.IDNumber = required.IDNumber
		d.PhoneNumber = required.PhoneNumber
		d.Realname = required.Realname
	}
	return
}

func (c *HospitalAdmin) UnmarshallJSON(data []byte) (err error) {
	required := struct {
		District string `json:"district" db:"district"`
		Hospital string `json:"business" db:"business"`
		Credit   string `json:"credit" db:"credit"`
		Address  string `json:"address" db:"address"`
		Head     string `json:"head" db:"hesd"`
		Phone    string `json:"phone" db:"phone"`
		ID       string `json:"ID" db:"ID"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Hospital) == 0 {
		err = errors.New("企业名称不能为空")
	} else if len(required.Credit) == 0 {
		err = errors.New("社会信用代码不能为空")
	} else if len(required.Address) == 0 {
		err = errors.New("注册地址不能为空")
	} else if len(required.Head) == 0 {
		err = errors.New("负责人不能为空")
	} else if len(required.Phone) == 0 {
		err = errors.New("负责人电话不能为空")
	} else if len(required.ID) == 0 {
		err = errors.New("负责人身份证不能为空")
	} else {
		c.District = required.District
		c.Hospital = required.Hospital
		c.Credit = required.Credit
		c.Address = required.Address
		c.Head = required.Head

		c.Phone = required.Phone
		c.ID = required.ID
	}
	return
}

func (c *Hospital) Unmarshall(data []byte) (err error) {
	required := struct {
		District string `json:"district" db:"district"`
		Hospital string `json:"business" db:"business"`
		Credit   string `json:"credit" db:"credit"`
		Address  string `json:"address" db:"address"`
		Head     string `json:"head" db:"hesd"`
		UserName string `json:"user_name" db:"user_name"`
		Phone    string `json:"phone" db:"phone"`
		ID       string `json:"ID" db:"ID"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Hospital) == 0 {
		err = errors.New("企业名称不能为空")
	} else if len(required.Credit) == 0 {
		err = errors.New("社会信用代码不能为空")
	} else if len(required.Address) == 0 {
		err = errors.New("注册地址不能为空")
	} else if len(required.Head) == 0 {
		err = errors.New("负责人不能为空")
	} else if len(required.Phone) == 0 {
		err = errors.New("负责人电话不能为空")
	} else if len(required.ID) == 0 {
		err = errors.New("负责人身份证不能为空")
	} else if len(required.UserName) == 0 {
		err = errors.New("用户名不能为空")
	} else {
		c.District = required.District
		c.Hospital = required.Hospital
		c.Credit = required.Credit
		c.Address = required.Address
		c.Head = required.Head
		c.UserName = required.UserName
		c.Phone = required.Phone
		c.ID = required.ID
	}
	return
}
