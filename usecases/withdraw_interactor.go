package usecases

import (
	"errors"
	"log"
	"strconv"
)

const EMPTY_STRUCT = 8

// 預金から残高を引き落とす
func (f FormTransactionCreditCustomer) Withdraw() error {
	var err error
	log.Println(f.CustomerId)

	selectCustomer, err := f.IsCustomerAndCredit()

	if err != nil {
		return errors.New("NO_CUSTOMER_ID")
	} else {
		err = selectCustomer.isValidWithdrawCredit(f)
		if err != nil {
			return err
		} else {
			var updateCreditBalance = strconv.Itoa(caliculateWithdrawCredit(selectCustomer.Credit_balance, f.TransactionCredit))

			err = selectCustomer.CustomerUpdate(updateCreditBalance)
			if err != nil {
				return errors.New("UPDATE_FAIL")

			}

			err := NewCreditHistory(f.CustomerId, f.TransactionCredit).RegisterTransacationHistory()
			if *err != nil {
				return errors.New("UPDATE_FAIL")
			}
		}
		return err
	}
}

//出勤時、入金に記録を残すためのクレジットID

//残高を引き落とす際に入力された値が有効か判定する
func (s Customer) isValidWithdrawCredit(formTransactionCustomer FormTransactionCreditCustomer) error {
	var err error
	transactionCredit, err := strconv.Atoi(formTransactionCustomer.TransactionCredit)
	if err != nil {
		return errors.New("INVALID_VALUE")

	}
	creditBalance, _ := strconv.Atoi(s.Credit_balance)
	if creditBalance < transactionCredit {
		return errors.New("NO_CASH")

	}
	return err
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

type WithdrawInterface interface {
	Withdraw() error
}
