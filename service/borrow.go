package serivce

import (
	"1earning/dao"
	"1earning/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BorrowBook(c *gin.Context) {
	id,_ := strconv.Atoi(c.Request.FormValue("id"))	
	userId := uint(id)
	book := model.Book{}
	book.Name = c.Request.FormValue("name")
	data :=dao.FindBookByName(book.Name)
	fmt.Println(data)
	
	if len(data) != 0 && data[0].Remaining > 0 {
		 ok := dao.BorrowBook(&book,userId)
		if ok {
			c.JSON(200,gin.H{
			"code":0,
			"message":"借阅成功",
			"data":data,
		})	
		return
		}
		c.JSON(200,gin.H{
			"code":-1,
			"message":"借阅失败",
			"data":-1,
		})
	}else{
		c.JSON(200,gin.H{
			"code":-1,
			"message":"借阅失败",
			"data":-1,
		})	
	}
}

func ReturnBook(c *gin.Context){
	book := model.Book{}
	book.Name = c.Request.FormValue("name")	
	id,_ := strconv.Atoi(c.Request.FormValue("id"))	
	userId := uint(id)
	ok := dao.ReturnBook(&book,userId)
	if ok {
		c.JSON(200,gin.H{
			"code":0,
			"message":"归还成功",
			"data":ok,


		})
	}else{
		c.JSON(200,gin.H{
			"code":0,
			"message":"归还失败",
			"data":ok,
		})
	}
	
}

//查询借书情况