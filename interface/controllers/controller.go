package controllers

import (
	"go_bank/entity"
	"go_bank/usecases"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//顧客を作成する際に値を受け取るための構造体
type RegisterCustomer struct {
	BranchNumer string `json:"branchNumber" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

//顧客の取引を行う際に受け取るための構造体
type TransactionCustomer struct {
	CustomerId         string `json:"customerId" validate:"required"`
	TrainsactionCredit string `json:"transactionCredit" validate:"required"`
}

type FormTransactionCustomer struct {
	*TransactionCustomer
}

type FormRegisterCustomer struct {
	*RegisterCustomer
}

type ValidationTransactionInterface interface {
	TransactionValidation() *entity.Credit_history
}

type ValidationCustomerInterface interface {
	Customer_validation() *entity.Customer
}

// func ValidateAll(e ValidationInterface) *entity.Credit_history {
// 	return e.WithDrawValidation() // インターフェースから呼び出す

// }

func Register(c *gin.Context) {
	var formRegisterCustomer RegisterCustomer
	var validation ValidationCustomerInterface

	if err := c.ShouldBindJSON(&formRegisterCustomer); err != nil {
		log.Fatal(err)
	} else {
		validate := validator.New()
		if err := validate.Struct(formRegisterCustomer); err != nil {
			log.Fatal(err)
		} else {
			log.Println(&formRegisterCustomer)
			validation = formRegisterCustomer
			usecases.Register(validation.Customer_validation())
			c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "Success"})
			log.Println(err)
		}

	}

}

func (t RegisterCustomer) Customer_validation() *entity.Customer {

	return &entity.Customer{
		Name:          t.Name,
		Branch_number: t.BranchNumer,
	}

}

func Withdraw(c *gin.Context) {
	var formTransactionCustomer TransactionCustomer
	var validation ValidationTransactionInterface
	// var sub = c.Request.Header.Get("sub")
	// log.Println(sub)
	if err := c.ShouldBindJSON(&formTransactionCustomer); err != nil {
		log.Fatal(err)
	}
	validate := validator.New()
	if err := validate.Struct(formTransactionCustomer); err != nil {
		log.Fatal(err)
	}

	validation = formTransactionCustomer
	var res = usecases.Withdraw(validation.TransactionValidation())
	if res.Result == true {
		c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "出金が完了しました。"})
	} else {
		c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: IndicateErrorMessage(res)})
	}

}

func (t TransactionCustomer) TransactionValidation() *entity.Credit_history {
	return &entity.Credit_history{
		Customer_id:        t.CustomerId,
		Transaction_credit: t.TrainsactionCredit,
	}
}

func Inquiry(c *gin.Context) {
	var formTransactionCustomer TransactionCustomer

	err := c.ShouldBindJSON(&formTransactionCustomer)

	c.JSON(http.StatusOK, entity.ReturnCredit{ResultCrecit: usecases.Inquiry(formTransactionCustomer.CustomerId)})

	if err != nil {
		println(err)
	}
}

func Deposit(c *gin.Context) {
	var formTransactionCustomer TransactionCustomer
	var validation ValidationTransactionInterface
	// var sub = c.Request.Header.Get("sub")
	// log.Println(sub)
	if err := c.ShouldBindJSON(&formTransactionCustomer); err != nil {
		log.Fatal(err)
	}

	validate := validator.New()
	if err := validate.Struct(formTransactionCustomer); err != nil {
		log.Fatal(err)
	}

	validation = formTransactionCustomer
	var res = usecases.Deposit(validation.TransactionValidation())
	if res.Result == true {
		c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "入金が完了しました。"})
	} else {
		c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: IndicateErrorMessage(res)})

	}

}

func IndicateErrorMessage(res *entity.ResultMessage) string {
	switch res.MessageType {
	case entity.NO_CUSTOMER_ID:
		return "該当の顧客IDが見つかりませんでした。"
	case entity.INVALID_VALUE:
		return "入力した値が不正です。"
	case entity.UPDATE_FAIL:
		return "再度時間を置いて実施してください。"
	case entity.NO_CASH:
		return "口座残高が不足しています。"
	}
	return "不正な処理が行われました。"
}
