package serivce

import (
	"1earning/dao"
	"1earning/model"
	"1earning/utils"
	"fmt"
	"math/rand"
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

//添加用户(注册用户)
func CreateUser(c *gin.Context) {
	user := model.UserBasic{}
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Phone = c.Request.FormValue("phone")
	user.Name = c.Request.FormValue("name")
	user.PassWord = c.Request.FormValue("password")
	repassword:= c.Request.FormValue("repassword")
	if repassword == user.PassWord {
		user.PassWord = utils.MakePassword(repassword, salt)
		user.Salt = salt
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
	if user.PassWord == repassword && id != 0{
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
	if data.Phone != ""{
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

//通过手机号和密码查找用户(登录)
func FindUserByPhoneAndPwd(c *gin.Context) {
	user := model.UserBasic{}
	phone := c.Request.FormValue("phone")
	passWord := c.Request.FormValue("password")
	user = dao.FindUserByPhone(phone)	
	
	if user.Phone == "" {
		c.JSON(200,gin.H{
			"code":-1,
			"message":"没有该用户",
			"data":-1,
		})
		return
	}
	fmt.Println(user.PassWord)
	flag := utils.ValidPassword(passWord,user.Salt,user.PassWord)	
	if !flag {
		c.JSON(200,gin.H{
			"code":-1,
			"message":"密码错误",
			"data":-1,
		})
		return
	}
	pwd := utils.MakePassword(passWord,user.Salt)
	data := dao.FindUserByPhoneAndPwd(user.Name,pwd)
	token := model.Token{}
	token.UserId = user.ID
	token.Value =  model.GenerateRandomString(32)
	dao.CreateToken(token)
	c.JSON(200,gin.H{
		"code":0,
		"message":"找到该用户，登录成功",
		"data":data,
	})
	
}