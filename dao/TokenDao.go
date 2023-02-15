package dao

import (
	"1earning/model"
	"1earning/utils"

	"gorm.io/gorm"
)

//创建Token

func CreateToken(token model.Token) *gorm.DB{
	return utils.DB.Create(&token)
}
