package usecases

import (
	"go_bank/entity"
	"math/rand"
	"strconv"
	"time"
)

const CREDIT_ID_LENGTH = 8 //頭にbranch_nameを足すために6桁にしている
const EMPTY_STRUCT = 8

// 預金から残高を引き落とす
func Withdraw(creditHistoryForm *entity.Credit_history) *entity.ResultMessage {
	var err error

	var selectCustomer = IsCustomerAndCredit(creditHistoryForm)

	if selectCustomer.Customer_id == "" {
		return &entity.ResultMessage{Result: false, MessageType: entity.NO_CUSTOMER_ID}
	} else {
		var res = isValidWithdrawCredit(selectCustomer.Credit_balance, creditHistoryForm)
		if res.Result == false {
			return res
		} else {
			var updateCreditBalance = strconv.Itoa(caliculateWithdrawCredit(selectCustomer.Credit_balance, creditHistoryForm.Transaction_credit))

			err = CustomerUpdate(selectCustomer.Customer_id, updateCreditBalance)
			if err != nil {
				return &entity.ResultMessage{Result: false, MessageType: entity.UPDATE_FAIL}
			}

			RegisterTransacationHistory(setCreditHistory(creditHistoryForm.Customer_id, creditHistoryForm.Transaction_credit))
		}
		return &entity.ResultMessage{Result: true, MessageType: entity.WITHDRAW_OK}
	}
}

//引き落とし時に記録を残すためのクレジットID
func generateCreditId() string {
	var creditId = ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < CREDIT_ID_LENGTH; i++ {
		creditId = creditId + strconv.Itoa(rand.Intn(10))
	}
	return creditId
}

//残高を引き落とす際に入力された値が有効か判定する
func isValidWithdrawCredit(selectCustomerCreditBalance string, creditHistoryForm *entity.Credit_history) *entity.ResultMessage {
	transactionCredit, err := strconv.Atoi(creditHistoryForm.Transaction_credit)
	if err != nil {
		return &entity.ResultMessage{Result: false, MessageType: entity.INVALID_VALUE}
	}
	creditBalance, _ := strconv.Atoi(selectCustomerCreditBalance)
	if creditBalance < transactionCredit {
		return &entity.ResultMessage{Result: false, MessageType: entity.NO_CASH}
	}
	return &entity.ResultMessage{Result: true, MessageType: entity.WITHDRAW_OK}
}

//預金から残高を引き落とす
func caliculateWithdrawCredit(selectCredit, formTransactionCredit string) int {
	intSelectCredit, _ := strconv.Atoi(selectCredit)
	intFormCreditBalance, _ := strconv.Atoi(formTransactionCredit)
	return intSelectCredit - intFormCreditBalance
}

func setCreditHistory(CustomerID, TransactionCredit string) *entity.Credit_history {

	return &entity.Credit_history{
		Customer_id:        CustomerID,
		Credit_id:          generateCreditId(),
		Transaction_credit: TransactionCredit,
		Credit_flag:        "1",
	}
}
