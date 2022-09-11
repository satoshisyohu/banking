package usecases

import (
	"errors"
	"go_bank/entity"

	"strconv"
)

// 預金から残高を引き落とす
func Deposit(formTransactionCustomer *entity.FormTransactionCustomer) error {
	var err error

	//container/listっていうライブラリを使うこともできるよ
	// l := list.New()
	// l.PushBack("z")
	// log.Println(l.Front().Value)

	selectCustomer, err := IsCustomerAndCredit(formTransactionCustomer)
	if &selectCustomer.Customer_id == nil {
		return errors.New("NO_CUSTOMER_ID")
	} else {
		err = selectCustomer.isValidDepositCredit(formTransactionCustomer)
		if err != nil {
			return err
		} else {
			var updateCreditBalance = strconv.Itoa(caliculateDepositCredit(selectCustomer.Credit_balance, formTransactionCustomer.TransactionCredit))

			err = selectCustomer.CustomerUpdate(updateCreditBalance)
			if err != nil {
				return errors.New("UPDATE_FAIL")
			}
			creditHistoryInterface := setcreditHistory(formTransactionCustomer.CustomerId, formTransactionCustomer.TransactionCredit)
			err := creditHistoryInterface.RegisterTransacationHistory()
			if err != nil {
				return *err
			}
		}
		return err
	}
}
func (c *Customer) isValidDepositCredit(formTransactionCustomer *entity.FormTransactionCustomer) error {
	_, err := strconv.Atoi(formTransactionCustomer.TransactionCredit)
	if err != nil {
		return errors.New("INVALID_VALUE")
		// return &entity.ResultMessage{Result: false, MessageType: entity.INVALID_VALUE}
	}
	_, err = strconv.Atoi(c.Credit_balance)
	if err != nil {
		return errors.New("INVALID_VALUE")
	}
	return err

}

//預金から残高を引き落とす
func caliculateDepositCredit(selectCredit, formTransactionCredit string) int {
	intSelectCredit, _ := strconv.Atoi(selectCredit)
	intFormCreditBalance, _ := strconv.Atoi(formTransactionCredit)
	return intSelectCredit + intFormCreditBalance
}

func setcreditHistory(CustomerID, TransactionCredit string) CreditHistoryInterface {

	return &CreditHistory{
		Customer_id:        CustomerID,
		Credit_id:          GenerateCreditId(),
		Transaction_credit: TransactionCredit,
		Credit_flag:        "0",
	}
}
