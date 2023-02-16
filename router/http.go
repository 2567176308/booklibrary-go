package router

import (
	serivce "1earning/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	//书籍
	// r.GET("/",serivce.GetBookList)
	//首页
	r.GET("/",serivce.GetIndex)
	//注册
	r.GET("/register",serivce.Register)
	r.GET("/login",serivce.LoginTo)
	r.GET("/booklibrary",serivce.GetBookList)
	r.Static("/layuiadmin","layuiadmin/")
	r.GET("book/getBookList",serivce.GetBookList)
	r.POST("book/createBook",serivce.CreateBook)
	r.POST("book/deleteBookByName",serivce.DeleteBookByName)
	r.POST("book/deleteBookByAuthor",serivce.DeleteBokByAuthor)
	r.POST("book/updateBook",serivce.UpdateBook)
	r.POST("book/findBookByName",serivce.FindBookByName)
	r.POST("book/findBookById",serivce.FindBookById)	
	//用户
	r.GET("user/getUserList",serivce.GetUserList)	
	/*注册*/
	r.POST("user/createUser",serivce.CreateUser)
	r.POST("user/deleteUser",serivce.DeleteUser)
	r.POST("user/updateUser",serivce.UpdateUser)
	r.POST("user/findUserByPhone",serivce.FindUserByPhone)
	/*登录*/
	r.POST("user/findUserByPhoneAndPwd",serivce.FindUserByPhoneAndPwd)
	//借阅
	r.POST("contact/borrowBook",serivce.BorrowBook)
	//归还
	r.POST("contact/returnBook",serivce.ReturnBook)
	return r
}