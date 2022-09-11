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

// type FormTransactionCustomer struct {
// 	*entity.FormTransactionCustomer
// }

// type FormCustomer struct {
// 	*entity.FormCustomer
// }

// type ValidationTransactionInterface interface {
// 	TransactionValidation() *entity.Credit_history
// }

// type ValidationCustomerInterface interface {
// 	Customer_validation() *entity.Customer
// }

type result entity.Result

func Register(c *gin.Context) {

	var formCustomer entity.FormCustomer
	// var validation ValidationCustomerInterface

	if err := c.ShouldBindJSON(&formCustomer); err != nil {
		log.Fatal(err)
	} else {
		validate := validator.New()
		if err := validate.Struct(formCustomer); err != nil {
			log.Fatal(err)
		} else {
			// validation = formFormCustomer
			// err := usecases.Register(validation.Customer_validation())
			err := usecases.Register(&formCustomer)

			if err != nil {
				log.Fatal(err)
			} else {
				c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "Success"})
			}
		}

	}

}

// func (t FormCustomer) Customer_validation() *entity.Customer {
// 	return &entity.Customer{
// 		Name:          t.Name,
// 		Branch_number: t.BranchNumer,
// 	}

func Withdraw(c *gin.Context) {
	var formTransactionCustomer entity.FormTransactionCustomer
	// var validation ValidationTransactionInterface
	// var sub = c.Request.Header.Get("sub")
	// log.Println(sub)
	if err := c.ShouldBindJSON(&formTransactionCustomer); err != nil {
		log.Fatal(err)
	}
	validate := validator.New()
	if err := validate.Struct(formTransactionCustomer); err != nil {
		log.Fatal(err)
	}

	// validation = formTransactionCustomer
	// var res = usecases.Withdraw(validation.TransactionValidation())
	log.Println(&formTransactionCustomer)
	res := usecases.Withdraw(&formTransactionCustomer)
	if res.Result == true {
		c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "出金が完了しました。"})
	} else {
		// c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: IndicateErrorMessage(res)})
	}

}

// func (t FormTransactionCustomer) TransactionValidation() *entity.Credit_history {
// 	return &entity.Credit_history{
// 		Customer_id:        t.CustomerId,
// 		Transaction_credit: t.TransactionCredit,
// 	}
// }

func Inquiry(c *gin.Context) {
	var formCustomer entity.FormInquieryCustomer

	if err := c.ShouldBindJSON(&formCustomer); err != nil {
		log.Fatal(err)
	}
	validate := validator.New()
	if err := validate.Struct(formCustomer); err != nil {
		log.Fatal(err)
	} else {
		creditBalance, err := usecases.Inquiry(formCustomer.CustomerId)
		if *err != nil {
			log.Println(*err)
		} else {
			c.JSON(http.StatusOK, entity.ReturnCredit{ResultCrecit: creditBalance})
		}
	}

}

func Deposit(c *gin.Context) {
	var formTransactionCustomer entity.FormTransactionCustomer
	// var validation ValidationTransactionInterface
	// var sub = c.Request.Header.Get("sub")
	// log.Println(sub)
	if err := c.ShouldBindJSON(&formTransactionCustomer); err != nil {
		log.Fatal(err)
	}

	validate := validator.New()
	if err := validate.Struct(formTransactionCustomer); err != nil {
		log.Fatal(err)
	}

	// validation = formTransactionCustomer
	// err := usecases.Deposit(validation.TransactionValidation())
	err := usecases.Deposit(&formTransactionCustomer)

	log.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: IndicateErrorMessage(err)})
	} else {
		c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "入金が完了しました。"})
	}

}

//これはどっか外に出した方が良い気がする
func IndicateErrorMessage(res error) string {
	print(res)
	switch res.Error() {
	case "NO_CUSTOMER_ID":
		return "該当の顧客IDが見つかりませんでした。"
	case "INVALID_VALUE":
		return "入力した値が不正です。"
	case "UPDATE_FAIL":
		return "再度時間を置いて実施してください。"
	case "NO_CASH":
		return "口座残高が不足しています。"
	}
	return "不正な処理が行われました。"
}
