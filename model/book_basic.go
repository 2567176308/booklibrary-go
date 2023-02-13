package model

import (
	"gorm.io/gorm"
)

type  Book struct {
	gorm.Model
	Name string // 书名
	Author string //作者
	Imag string //封面图片
	Price int //价格
	Press string //出版社
	Type string //类型
	Reserve int //预约数量
	Remaining int //剩余本数
	Loan int //借出数量

}

func (table *Book) TableName() string {
	return "book"
}
