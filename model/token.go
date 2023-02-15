package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserId uint
	Value string
	ExpiresAt time.Time
}

func (table Token) TableName() string {
	return "token"
}


func (t *Token) BeforeCreate(tx *gorm.DB) error {
    t.ExpiresAt = time.Now().Add(time.Hour) // 设置token过期时间为1小时
    return nil
}
