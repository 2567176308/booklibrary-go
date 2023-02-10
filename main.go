package main

import (
	"1earning/router"
	"1earning/utils"
	// "github.com/gin-gonic/gin"
)

func main() {
    r := router.Router()
	utils.InitConfig()
	utils.InitMySQL()
    r.Run() // listen and serve on 0.0.0.0:8080
}

