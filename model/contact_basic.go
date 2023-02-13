package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	UserId uint
	BookId uint
	Desc string //描述，备注	
}

func (table Contact)TableName() string {
	return "contact"
}