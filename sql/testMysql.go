/**
* @Auth:ShenZ
* @Description:
* @CreateDate:2022/06/15 10:57:44
 */
package main

import (
	"1earning/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:00000.@tcp(127.0.0.1:3306)/booklibrary?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	// db.AutoMigrate(&models.Community{})
	// db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Contact{})
	// db.AutoMigrate(&models.Message{})
	// db.AutoMigrate(&models.GroupBasic{})
	// db.AutoMigrate(&models.Contact{})

	// Create
	// user := &models.UserBasic{}
	// user.Name = "申专"
	// db.Create(user)

	// // Read
	// fmt.Println(db.First(user, 1)) // 根据整型主键查找
	// //db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// // Update - 将 product 的 price 更新为 200
	// db.Model(user).Update("PassWord", "1234")
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
