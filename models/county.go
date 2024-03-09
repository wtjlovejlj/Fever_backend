package models

import (
	"encoding/json"
	"errors"
	"time"
)

type County struct {
	Card      int64     `json:"card" db:"card" gorm:"primaryKey"` // 编号
	District  string    `json:"district" db:"district"`           // 区县名称
	UserName  string    `json:"username" db:"username"`           // 用户名
	UpdatedAt time.Time `db:"updated_at"`
}

func (con *County) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		District string `json:"district" db:"district"` // 区县名称
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.District) == 0 {
		err = errors.New("缺少必填字段District")
	} else {
		con.District = required.District
	}
	return
}
