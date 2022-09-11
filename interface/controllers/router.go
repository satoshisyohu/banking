package controllers

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {

	r := gin.Default()
	r.POST("/register", Register) //顧客登録
	r.POST("/withdraw", Withdraw) //出金
	r.POST("/deposit", Deposit)   //入金
	r.POST("/inquiry", Inquiry)   //残高確認
	// r.POST("/interst/", Interest)

	return r
}
