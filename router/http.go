package router

import (
	serivce "1earning/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/",serivce.GetBookList)
	r.GET("book/getBookList",serivce.GetBookList)
	r.POST("book/createBook",serivce.CreateBook)
	r.POST("book/deleteBookByName",serivce.DeleteBookByName)
	r.POST("book/deleteBookByAuthor",serivce.DeleteBokByAuthor)
	r.POST("book/updateBook",serivce.UpdateBook)
	r.POST("book/findBookByName",serivce.FindBookByName)
	r.POST("book/findBookById",serivce.FindBookById)	
	return r
}