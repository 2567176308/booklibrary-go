package router

import (
	serivce "1earning/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	//书籍
	// r.GET("/",serivce.GetBookList)
	r.GET("book/getBookList",serivce.GetBookList)
	r.POST("book/createBook",serivce.CreateBook)
	r.POST("book/deleteBookByName",serivce.DeleteBookByName)
	r.POST("book/deleteBookByAuthor",serivce.DeleteBokByAuthor)
	r.POST("book/updateBook",serivce.UpdateBook)
	r.POST("book/findBookByName",serivce.FindBookByName)
	r.POST("book/findBookById",serivce.FindBookById)	
	//用户
	r.GET("user/getUserList",serivce.GetUserList)	
	r.POST("user/createUser",serivce.CreateUser)
	r.POST("user/deleteUser",serivce.DeleteUser)
	r.POST("user/updateUser",serivce.UpdateUser)
	r.POST("user/findUserByPhone",serivce.FindUserByPhone)
	//借阅
	r.POST("contact/borrowBook",serivce.BorrowBook)
	//归还
	r.POST("contact/returnBook",serivce.ReturnBook)
	return r
}