package serivce

import (
	"1earning/dao"
	"1earning/model"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/hashicorp/hcl/hcl/strconv"
)

//显示书籍列表
func GetBookList(c *gin.Context) { 
	data := dao.GetBookList()
	c.JSON(200,gin.H{
		"code": 0,//0成功,1失败
		"message":"查询成功",
		"data":data,	
	})
}
//新增书籍
func CreateBook(c *gin.Context) {
	book := model.Book{}
	book.Name = c.Request.FormValue("name")
	book.Author = c.Request.FormValue("author")
	price := c.Request.FormValue("price")
	book.Press = c.Request.FormValue("press")
	book.Type = c.Request.FormValue("type")
	book.Price,_ = strconv.Atoi(price)
	dao.CreateBook(book)
	c.JSON(200,gin.H{
		"code": 0,//0成功,1失败
		"message":"书籍新增成功",
		"data":book,	
	})

}
//通过书名删除书籍
func DeleteBookByName(c *gin.Context) {
	book := model.Book{}
	book.Name = c.Request.FormValue("name")
	data := dao.FindBookByName(book.Name)
	//判断是否存在该书籍名称
	if len(data) == 0{
		c.JSON(200,gin.H{
			"code":-1,
			"message":"书籍删除失败",
			"data": book.Name,
		})
	return
	}else{
		dao.DeleteBookByName(book)
			c.JSON(200,gin.H{
			"code":0,
			"message":"书籍删除成功",
			"data": book.Name,
		})
	}
	

//通过作者名删除书籍
/*...*/

}
func DeleteBokByAuthor(c *gin.Context){
	book := model.Book{}
	book.Author = c.Request.FormValue("author")
	if len(dao.FindBookByAuthor(book.Author)) != 0{
	dao.DeleteBokByAuthor(book)	
	c.JSON(200,gin.H{
	"code":0,
	"message":"书籍删除成功",
	"data": book.Author,
		})
	return	
	}else{
		c.JSON(200,gin.H{
		"code":-1,
		"message":"书籍删除失败",//没有该作者的书籍
		"data": book.Author,
			})
	}
}

//修改书籍信息(包括书名，作者，出版社信息,价格，封面图片路径)
func UpdateBook(c *gin.Context){
	book := model.Book{}
	id,_ := strconv.Atoi(c.Request.FormValue("id"))
	book.ID = uint(id)
	book.Name = c.Request.FormValue("name")
	book.Author = c.Request.FormValue("author")
	book.Press = c.Request.FormValue("press")
	price,_ := strconv.Atoi(c.Request.FormValue("price"))
	book.Type = c.Request.FormValue("type")
	book.Price = price
	// book.Imag = c.Request.FormValue("imag")
	if len(dao.FindBookById(book.ID)) != 0{
		dao.UpdateBook(book)	
			c.JSON(200,gin.H{
				"code":0,
				"message":"书籍信息修改成功",
				"data":book,
			})

	}else{
		c.JSON(200,gin.H{
				"code":-1,
				"message":"没有该书籍",
				"data":-1,
			})
	}
	

}
//查询相应书籍

func FindBookByName(c *gin.Context) {
	book := model.Book{}
	book.Name = c.Request.FormValue("name")
	data := dao.FindBookByName(book.Name)
	if len(data) != 0{
		c.JSON(200,gin.H{
			"code":0,
			"message":"查询到含有该书籍",
			"data":data,
		})
	}else{
		c.JSON(200,gin.H{
		"code":-1,
		"message":"没有该书籍",
		"data":-1,
	})
	}
}
//通过id查找书籍
func FindBookById(c *gin.Context) {
	book := model.Book{}
	id,_ := strconv.Atoi(c.Request.FormValue("id"))
	book.ID = uint(id)
	data := dao.FindBookById(book.ID)
	if len(data) != 0 {
		c.JSON(200,gin.H{
			"code":0,
			"message":"找到该书籍",
			"data":data,
		})
	}else{
		c.JSON(200,gin.H{
			"code":-1,
			"message":"没有该书籍",
			"data":-1,
		})	
	}

}



