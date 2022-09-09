package controllers

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {

	r := gin.Default()
	r.POST("/register", Register) //顧客登録
	r.POST("/withdraw", Withdraw) //出金
	r.POST("/deposit", Deposit)   //入金
	r.POST("/inquiry", Inquiry)   //出金
	//最終確認
	r.POST("/revert", Inquiry)
	//最終確認
	return r
}
