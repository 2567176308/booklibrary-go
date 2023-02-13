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
		dao.BorrowBook(&book,userId)

		c.JSON(200,gin.H{
			"code":0,
			"message":"借阅成功",
			"data":data,
		})	

		
	}else{
		c.JSON(200,gin.H{
			"code":-1,
			"message":"借阅失败",
			"data":-1,
		})	
	}
}