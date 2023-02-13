package dao

import (
	"1earning/model"
	"1earning/utils"
	"fmt"
)

func GetContactList() []*model.Contact {
	data := make([]*model.Contact, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func BorrowBook(book *model.Book,userId uint) {
	contact := model.Contact{}
	contact.UserId = userId
	utils.DB.Where("name = ?", book.Name).First(&book)
	book.Remaining = book.Remaining - 1
	book.Loan = book.Loan + 1
	contact.BookId = book.ID
	utils.DB.Save(&book)
	utils.DB.Create(&contact)
	fmt.Println("[borrow]name>>>", book.Name)
	fmt.Println("[borrow]UserId>>>",userId)
}