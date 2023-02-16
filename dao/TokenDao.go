package dao

import (
	"1earning/model"
	"1earning/utils"
	"time"
	"gorm.io/gorm"
)

//创建Token

func CreateToken(token model.Token) *gorm.DB{
	return utils.DB.Create(&token)
}

//查询token是否已经过期
func TokenScope(token model.Token,userId uint) *gorm.DB {
    return utils.DB.Where("expires_at > ? and user_id", time.Now(),userId).First(&token)
}