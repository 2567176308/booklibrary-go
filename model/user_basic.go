package model

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name string
	Phone string
	PassWord string
	Avatar string //头像
	DeviceInfo []Book //保留关键字
	Salt string //加密字符串


}
func (table *UserBasic) TableName() string {
	return "user_basic"
}