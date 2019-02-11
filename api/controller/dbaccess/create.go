package dbaccess

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"simple_rest/database"
	"simple_rest/env"

	"github.com/gin-gonic/gin"
)

type UserDataInput struct{
	Account string `form:"Account"`
	Password string `form:"Password"`
}

type ResultOutput struct{
	IsOK bool
}

func Create(c *gin.Context){
	res := protocol.Response{}
	input := &UserDataInput{}
	r := ResultOutput{IsOK:false}
	res.Result = &r

	if err := c.Bind(input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, r)
		return
	}

	if err:= InsertUser(input); err != nil {
		c.JSON(http.StatusInternalServerError, protocol.SomethingWrongRes(err))
		return
	}

	r.IsOK = true
	c.JSON(http.StatusOK, res)
	return
}

func InsertUser(data *UserDataInput) (err error){
	fn := "InsertUser"

	dbS := database.GetConn(env.AccountDB)

	sql := " INSERT INTO "
	sql += " user "
	sql += " (account, password) "
	sql += " VALUES "
	sql += " (?,?) ;"

	_, err = dbS.Exec(sql, data.Account, data.Password)
	if err != nil {
		log.Fatalf("Exec Insert Failed. fn:%s , error:%s", fn, err.Error())
		return
	}

	return
}