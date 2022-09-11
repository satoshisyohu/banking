package usecases

import (
	"go_bank/entity"
	"log"
	"strconv"
)

const EMPTY_STRUCT = 8

// 預金から残高を引き落とす
func Withdraw(formTransactionCustomer *entity.FormTransactionCustomer) *entity.ResultMessage {
	var err error
	log.Println(formTransactionCustomer.CustomerId)

	selectCustomer, err := IsCustomerAndCredit(formTransactionCustomer)

	if err != nil {
		return &entity.ResultMessage{Result: false, MessageType: entity.NO_CUSTOMER_ID}
	} else {
		var res = selectCustomer.isValidWithdrawCredit(formTransactionCustomer)
		if res.Result == false {
			return res
		} else {
			var updateCreditBalance = strconv.Itoa(caliculateWithdrawCredit(selectCustomer.Credit_balance, formTransactionCustomer.TransactionCredit))

			err = selectCustomer.CustomerUpdate(updateCreditBalance)
			if err != nil {
				return &entity.ResultMessage{Result: false, MessageType: entity.UPDATE_FAIL}
			}

			err := NewCreditHistory(formTransactionCustomer.CustomerId, formTransactionCustomer.TransactionCredit).RegisterTransacationHistory()
			if *err != nil {
				return &entity.ResultMessage{Result: false, MessageType: entity.UPDATE_FAIL}
			}
		}
		return &entity.ResultMessage{Result: true, MessageType: entity.WITHDRAW_OK}
	}
}

//出勤時、入金に記録を残すためのクレジットID

//残高を引き落とす際に入力された値が有効か判定する
func (s Customer) isValidWithdrawCredit(formTransactionCustomer *entity.FormTransactionCustomer) *entity.ResultMessage {
	transactionCredit, err := strconv.Atoi(formTransactionCustomer.TransactionCredit)
	if err != nil {
		return &entity.ResultMessage{Result: false, MessageType: entity.INVALID_VALUE}
	}
	creditBalance, _ := strconv.Atoi(s.Credit_balance)
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

func NewCreditHistory(CustomerID, TransactionCredit string) *CreditHistory {

	return &CreditHistory{
		Customer_id:        CustomerID,
		Credit_id:          GenerateCreditId(),
		Transaction_credit: TransactionCredit,
		Credit_flag:        "1",
	}
}
