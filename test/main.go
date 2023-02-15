package main

import (
	"1earning/dao"
	"1earning/model"

)

func main() {
	token := model.Token{}
	token.Value = "1293"
	dao.CreateToken(token)
}