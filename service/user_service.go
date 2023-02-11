package serivce

import (
	"1earning/dao"
	"1earning/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := dao.GetUserList()
	c.JSON(200, gin.H{
		"code":    0, //0成功,1失败
		"message": "查询成功",
		"data":    data,
	})
}

//添加用户
func CreateUser(c *gin.Context) {
	user := model.UserBasic{}
	user.Phone = c.Request.FormValue("phone")
	user.Name = c.Request.FormValue("name")
	user.PassWord = c.Request.FormValue("password")
	repassword:= c.Request.FormValue("repassword")
	if repassword == user.PassWord {
		dao.CreateUser(user)
		c.JSON(200,gin.H{
			"code":0,
			"message":"用户注册成功",
			"data":user,
		})
		return
	}else{
		c.JSON(200,gin.H{
			"code":-1,
			"message":"用户注册失败",
			"data":-1,
		})
	}
}

//删除用户
func DeleteUser(c *gin.Context) {
	user := model.UserBasic{}
	id,_ := strconv.Atoi(c.Request.FormValue("id"))
	user.ID = uint(id)
	dao.DeleteUser(user)
	c.JSON(200,gin.H{
			"code":0,
			"message":"用户删除成功",
			"data":id,
		})
}

//修改用户信息
func UpdateUser(c *gin.Context) {
	user := model.UserBasic{}	
	id,_ := strconv.Atoi(c.Request.FormValue("id"))
	user.ID = uint(id)
	user.Name = c.Request.FormValue("name")	
	user.Phone = c.Request.FormValue("phone")
	user.PassWord = c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")
	if user.PassWord == repassword {
		dao.UpdateUser(user)
		c.JSON(200,gin.H{
			"code":0,
			"message":"用户修改成功",
			"data":id,
		})
		
	}else{
		c.JSON(200,gin.H{
			"code":-1,
			"message":"用户修改失败",
			"data":-1,
		})
	}
	
}


//查找用户
func FindUserByPhone(c *gin.Context) {
	phone := c.Request.FormValue("phone")	
	//判断正误
	// 。。。
	data := dao.FindUserByPhone(phone)
	if len(data) != 0{
		c.JSON(200,gin.H{
			"code":0,
			"message":"查询到该用户",
			"data":data,
		})
	}else{
		c.JSON(200,gin.H{
			"code":0,
			"message":"没有该用户",
			"data":-1,
		})
	}

}