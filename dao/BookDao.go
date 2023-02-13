package dao

import (
	"1earning/model"
	"1earning/utils"
	"fmt"

	"gorm.io/gorm"
)


func GetBookList() []*model.Book {
	data := make([]*model.Book,10)
	utils.DB.Find(&data)
	for _,v := range data {
		fmt.Println(v)
	}
	return data	
}




func CreateBook(book model.Book) *gorm.DB{
	return utils.DB.Create(&book)
}

func DeleteBookByName(book model.Book) *gorm.DB {
	fmt.Println("[delete]name>>>",book.Name)
	return utils.DB.Where("name = ?",book.Name).Delete(&book)
}

func DeleteBokByAuthor(book model.Book) *gorm.DB {
	fmt.Println("[delete]author>>>",book.Author)
	return utils.DB.Where("author = ?",book.Author).Delete(&book)
}


func UpdateBook(book model.Book) *gorm.DB {
	fmt.Println("[update]id>>>",book.ID)
	return utils.DB.Model(&book).Updates(model.Book{Name: book.Name,Author: book.Author,
		Press: book.Press,Price: book.Price,Type: book.Type,})
}
//通过作者查找书籍
func FindBookByAuthor(author string) []*model.Book {
	data := make([]*model.Book,10)
	utils.DB.Where("author = ?",author).Find(&data)
	fmt.Println("[find]author>>>",author)
	for _,v := range(data){
		fmt.Println(v)
	}
	return data	
}
//通过名字查找书籍
func FindBookByName(name string) []*model.Book {
	data := make([]*model.Book,10)
	utils.DB.Where("name = ?",name).Find(&data)
	fmt.Println("[find]name>>>",name)
	for _,v := range(data){
		fmt.Println(v)
	}
	return data
}
//通过id查找书籍
func FindBookById(id uint) []*model.Book {
	data := make([]*model.Book,10)
	utils.DB.Where("id = ?",id).First(&data)
	fmt.Println("[find]id>>>",id)
	return data
}

//借阅书籍


