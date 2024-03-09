package models

import (
	"encoding/json"
	"errors"
	"time"
)

type FeverPeople struct {
	FeverID          uint64    `json:"fever_id" db:"fever_id"`
	Address          string    `json:"address" db:"address"`                     // 现住址
	Age              int64     `json:"age" db:"age"`                             // 年龄
	ClinicalEffect   string    `json:"clinical_effect" db:"clinical_effect"`     // 临床表现
	ColorCode        bool      `json:"color_code" db:"color_code"`               // 红黄码检测
	Date             time.Time `db:"creat_at" db:"creat_at"`                     // 时间日期
	DisposalMethod   string    `json:"disposal_method" db:"disposal_method"`     // 处置方式
	Doctor           string    `json:"doctor" db:"doctor"`                       // 接诊医生
	IDCardID         string    `json:"id_card_id" db:"id_card_id"`               // 身份证号
	MobilePhone      string    `json:"mobile_phone" db:"mobile_phone"`           // 手机号
	Name             string    `json:"name" db:"name"`                           // 姓名
	NucleicAcidTest  bool      `json:"nucleic_acid_test" db:"nucleic_acid_test"` // 核酸检测
	PrimaryiDagnosis string    `json:"primaryi_dagnosis" db:"primaryi_dagnosis"` // 初步诊断
	ProgramResults   string    `json:"program_results" db:"program_results"`     // 检测项目结果
	Sex              string    `json:"sex" db:"sex"`                             // 性别
}

func (u *FeverPeople) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Address          string `json:"address"`           // 现住址
		Age              int64  `json:"age"`               // 年龄
		ClinicalEffect   string `json:"clinical_effect"`   // 临床表现
		ColorCode        bool   `json:"color_code"`        // 红黄码检测
		DisposalMethod   string `json:"disposal_method"`   // 处置方式
		Doctor           string `json:"doctor"`            // 接诊医生
		IDCardID         string `json:"id_card_id"`        // 身份证号
		MobilePhone      string `json:"mobile_phone"`      // 手机号
		Name             string `json:"name"`              // 姓名
		NucleicAcidTest  bool   `json:"nucleic_acid_test"` // 核酸检测
		PrimaryiDagnosis string `json:"primaryi_dagnosis"` // 初步诊断
		ProgramResults   string `json:"program_results"`   // 检测项目结果
		Sex              string `json:"sex"`               // 性别
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Address) == 0 {
		err = errors.New("缺少必填字段Address")
	} else if required.Age == 0 {
		err = errors.New("缺少必填字段Age")
	} else if len(required.ClinicalEffect) == 0 {
		err = errors.New("缺少必填字段ClinicalEffect")
	} else if len(required.DisposalMethod) == 0 {
		err = errors.New("缺少必填字段DisposalMethod")
	} else if len(required.Doctor) == 0 {
		err = errors.New("缺少必填字段Doctor")
	} else if len(required.IDCardID) != 18 {
		err = errors.New("缺少必填字段IDCardID")
	} else if len(required.MobilePhone) != 11 {
		err = errors.New("缺少必填字MobilePhone")
	} else if len(required.Name) == 0 {
		err = errors.New("缺少必填字段Name")
	} else if len(required.PrimaryiDagnosis) == 0 {
		err = errors.New("缺少必填字段PrimaryiDagnosis")
	} else if len(required.ProgramResults) == 0 {
		err = errors.New("缺少必填字段ProgramResults")
	} else if len(required.Sex) == 0 {
		err = errors.New("缺少必填字段Sex")
	} else {
		u.Address = required.Address
		u.Age = required.Age
		u.ClinicalEffect = required.ClinicalEffect
		u.ColorCode = required.ColorCode
		u.DisposalMethod = required.DisposalMethod
		u.Doctor = required.Doctor
		u.IDCardID = required.IDCardID
		u.MobilePhone = required.MobilePhone
		u.Name = required.Name
		u.NucleicAcidTest = required.NucleicAcidTest
		u.PrimaryiDagnosis = required.PrimaryiDagnosis
		u.ProgramResults = required.ProgramResults
		u.Sex = required.Sex
	}
	return
}
