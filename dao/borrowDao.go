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

func BorrowBook(book *model.Book,userId uint) bool {
	contact := model.Contact{}
	contact.UserId = userId
	utils.DB.Where("name = ?", book.Name).First(&book)
	book.Remaining = book.Remaining - 1
	book.Loan = book.Loan + 1
	contact.BookId = book.ID
	contact_data := FindContactFirst(userId,book.ID)
		if len(contact_data) == 0  {
		utils.DB.Save(&book)
		utils.DB.Create(&contact)
		fmt.Println("[borrow]name>>>", book.Name)
		fmt.Println("[borrow]UserId>>>",userId)

		return true
	}
	return false
	
	
}

func ReturnBook(book *model.Book,userId uint)  bool{
	contact := model.Contact{}
	utils.DB.Where("name = ?",book.Name).First(&book)
	fmt.Println("[return]book>>>",book)
	book.Remaining = book.Remaining + 1
	book.Loan = book.Loan -1
	contact_data := FindContactFirst(userId,book.ID)
	if len(contact_data) != 0{
		utils.DB.Save(&book)
		utils.DB.Where("user_id = ? and book_id = ?",userId,book.ID).Delete(&contact)
		return true
	}else{
		return false
	}
	
}

func FindContactFirst(userId,booId uint) []*model.Contact {
	data := make([]*model.Contact,10)
	utils.DB.Where("user_id = ? and book_id = ?",userId,booId).First(&data)
	return data
}


//查询借书情况
