package serivce

import (
	"fmt"
	"text/template"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

func Register(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/reg.html")
		if err != nil {
			fmt.Println(err)
		}
		ind.Execute(c.Writer, "register")
}

func LoginTo(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/login-2.html")
	if err != nil{
		fmt.Println(err)
	}
	ind.Execute(c.Writer,"login")
}