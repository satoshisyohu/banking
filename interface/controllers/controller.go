package controllers

import (
	"errors"
	"go_bank/entity"
	"go_bank/usecases"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type result entity.Result

type interfaceStruct struct {
	RegisterInterface usecases.RegisterInterface
	DepositInterface  usecases.DepositInterface
	WithdrawInterface usecases.WithdrawInterface
	InquiryInterface  usecases.InquiryInterface
	TransferInterface usecases.TransferInterface
}

func Register(c *gin.Context) {

	var formCustomer usecases.FormCusotmoer

	if err := c.ShouldBindJSON(&formCustomer); err != nil {
		log.Println(errors.New("JsonMappingに失敗しました。"))
	} else {
		validate := validator.New()
		if err := validate.Struct(formCustomer); err != nil {
			log.Println(errors.New("ValidationErrorが発生しました。"))
		} else {
			RegisterInterface := &formCustomer
			err := RegisterInterface.Register()

			if err != nil {
				log.Println(errors.New("顧客登録に失敗しました。"))
			} else {
				c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "Success"})
			}
		}

	}

}

func Deposit(c *gin.Context) {
	var formTransactionCustomer usecases.FormTransactionCreditCustomer

	if err := c.ShouldBindJSON(&formTransactionCustomer); err != nil {
		log.Println(formTransactionCustomer)
		log.Println(errors.New("JsonMappingに失敗しました。"))
	} else {
		validate := validator.New()
		if err := validate.Struct(formTransactionCustomer); err != nil {
			log.Println(errors.New("ValidationErrorが発生しました。"))
		}

		var DepositInterface interfaceStruct
		DepositInterface.DepositInterface = &formTransactionCustomer
		err := DepositInterface.DepositInterface.Deposit()

		if err != nil {
			c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: IndicateErrorMessage(err)})
		} else {
			c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "入金が完了しました。"})
		}
	}
}

func Withdraw(c *gin.Context) {
	var formTransactionCustomer usecases.FormTransactionCreditCustomer

	if err := c.ShouldBindJSON(&formTransactionCustomer); err != nil {
		log.Println(errors.New("JsonMappingに失敗しました。"))
	} else {
		validate := validator.New()
		if err := validate.Struct(formTransactionCustomer); err != nil {
			log.Println(errors.New("ValidationErrorが発生しました。"))
		}

		var WithdrawInterface interfaceStruct
		WithdrawInterface.WithdrawInterface = &formTransactionCustomer
		err := WithdrawInterface.WithdrawInterface.Withdraw()
		if err != nil {
			c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: IndicateErrorMessage(err)})
		} else {
			c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "出金が完了しました。"})

		}
	}
}

func Inquiry(c *gin.Context) {
	var formCustomer usecases.FormInquieryCustomer

	if err := c.ShouldBindJSON(&formCustomer); err != nil {
		log.Println(errors.New("JsonMappingに失敗しました。"))
	} else {
		validate := validator.New()
		if err := validate.Struct(formCustomer); err != nil {
			log.Println(errors.New("ValidationErrorが発生しました。"))
		} else {
			var InquiryInterface interfaceStruct
			InquiryInterface.InquiryInterface = &formCustomer
			creditBalance, err := InquiryInterface.InquiryInterface.Inquiry()

			if err != nil {
				c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: IndicateErrorMessage(err)})
			} else {
				c.JSON(http.StatusOK, entity.ReturnCredit{ResultCrecit: creditBalance})
			}
		}
	}
}

func Transfer(c *gin.Context) {
	var err error
	var formCustomerId usecases.FormInquieryCustomer
	formCustomerId.CustomerId = c.GetHeader("customer-id")
	var formTransferCustomer usecases.FormTransferCustomer
	log.Println(formTransferCustomer)
	if err = c.ShouldBindJSON(&formTransferCustomer); err != nil {
		log.Println(errors.New("JsonMappingに失敗しました。"))
	} else {
		var TransferInterface interfaceStruct
		TransferInterface.TransferInterface = &formTransferCustomer
		err = formTransferCustomer.Transfer(formCustomerId)
		if err != nil {
			c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: IndicateErrorMessage(err)})

		} else {
			c.JSON(http.StatusOK, entity.ReturnResult{ResultMessage: "振り込みが完了しました。"})
		}

		//headrerのチェック

		//jsonobjectのマッピングチェック

		//interactorに処理を渡す

		//jsonでreturnを返す

	}
}

func Gmap(c *gin.Context) {
	var err error
	// res, err := usecases.Gmap()
	if err != nil {
		c.HTML(200, "", "https://www.google.com/maps/search/?api=1&query=スターバックス")

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
