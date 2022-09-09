package usecases

import (
	"container/list"
	"go_bank/entity"
	"log"
	"strconv"
)

type CreditHistory entity.Credit_history

// 預金から残高を引き落とす
func Deposit(creditHistoryForm *entity.Credit_history) *entity.ResultMessage {
	var err error

	//container/listっていうライブラリを使うこともできるよ
	l := list.New()
	l.PushBack("z")
	log.Println(l.Front().Value)

	var selectCustomer = IsCustomerAndCredit(creditHistoryForm)

	if selectCustomer.Customer_id == "" {
		return &entity.ResultMessage{Result: false, MessageType: entity.NO_CUSTOMER_ID}
	} else {
		var creditHistory CreditHistory
		creditHistory = CreditHistory(*creditHistoryForm)
		var res = creditHistory.isValidDepositCredit(selectCustomer.Credit_balance)
		if res.Result == false {
			return res
		} else {
			var updateCreditBalance = strconv.Itoa(caliculateDepositCredit(selectCustomer.Credit_balance, creditHistoryForm.Transaction_credit))

			err = CustomerUpdate(selectCustomer.Customer_id, updateCreditBalance)
			if err != nil {
				return &entity.ResultMessage{Result: false, MessageType: entity.UPDATE_FAIL}
			}

			(RegisterTransacationHistory(setcreditHistory(creditHistoryForm.Customer_id, creditHistoryForm.Transaction_credit)))
		}
		return &entity.ResultMessage{Result: true, MessageType: entity.DEPOSIT_OK}
	}
}
func (c *CreditHistory) isValidDepositCredit(selectCustomerCreditBalance string) *entity.ResultMessage {
	_, err := strconv.Atoi(c.Transaction_credit)
	if err != nil {
		return &entity.ResultMessage{Result: false, MessageType: entity.INVALID_VALUE}
	}
	_, err = strconv.Atoi(selectCustomerCreditBalance)
	if err != nil {
		return &entity.ResultMessage{Result: false, MessageType: entity.INVALID_VALUE}
	}
	return &entity.ResultMessage{Result: true, MessageType: entity.DEPOSIT_OK}

}

//預金から残高を引き落とす
func caliculateDepositCredit(selectCredit, formTransactionCredit string) int {
	intSelectCredit, _ := strconv.Atoi(selectCredit)
	intFormCreditBalance, _ := strconv.Atoi(formTransactionCredit)
	return intSelectCredit + intFormCreditBalance
}

func setcreditHistory(CustomerID, TransactionCredit string) *entity.Credit_history {

	return &entity.Credit_history{
		Customer_id:        CustomerID,
		Credit_id:          generateCreditId(),
		Transaction_credit: TransactionCredit,
		Credit_flag:        "0",
	}
}
