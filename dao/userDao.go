package dao

import (
	"1earning/model"
	"1earning/utils"
	"fmt"

	"gorm.io/gorm"
)

func GetUserList() []*model.UserBasic {
	data := make([]*model.UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
//增
func CreateUser(user model.UserBasic) *gorm.DB{
	return utils.DB.Create(&user)
}
//删
func DeleteUser(user model.UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
//改
func UpdateUser(user model.UserBasic) *gorm.DB{
	return utils.DB.Updates(&user)
}
//查
func FindUserByPhone(phone string) []*model.UserBasic {
	data := make([]*model.UserBasic,10)
	utils.DB.Where("phone = ?",phone).Find(&data)
	return data	
}
