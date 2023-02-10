package main

import (
	"1earning/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  type Product struct {
	gorm.Model
	Code  string
	Price uint
  }
  
  func main() {
	db, err := gorm.Open(mysql.Open("root:00000.@tcp(127.0.0.1:3306)/booklibrary?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// 迁移 schema
	db.AutoMigrate(&model.Book{})
  
	// Create
	// db.Create(&model.Book{Name: "活着", Price: 100})
  
	// Read
	// var product Product
	// db.First(&product, 1) // 根据整型主键查找
	// db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
  
	// Update - 将 product 的 price 更新为 200
	// db.Model(&model.Book{}).Update("Price", 200)
	// Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
  
	// Delete - 删除 product
	// db.Delete(&product, 1)
  }