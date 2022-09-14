package usecases

import (
	"errors"
)

const EMPTY_STRUCT = 8

type WithdrawInterface interface {
	Withdraw() error
}

// 預金から残高を引き落とす
func (f FormTransactionCreditCustomer) Withdraw() error {
	var err error

	selectCustomer, err := f.IsCustomerAndCredit()

	if err != nil {
		return errors.New("NO_CUSTOMER_ID")
	} else {
		err = selectCustomer.isValidWithdrawCredit(f)
		if err != nil {
			return err
		} else {

			err = selectCustomer.CustomerUpdate()
			if err != nil {
				return errors.New("UPDATE_FAIL")

			}
			creditHistoryInterface := f.NewCreditHistory()
			err := creditHistoryInterface.RegisterTransacationHistory()
			if *err != nil {
				return errors.New("UPDATE_FAIL")
			}
		}
		return err
	}
}

//出勤時、入金に記録を残すためのクレジットID

//残高を引き落とす際に入力された値が有効か判定する
func (s *Customer) isValidWithdrawCredit(formTransactionCustomer FormTransactionCreditCustomer) error {
	var err error
	if s.Credit_balance < formTransactionCustomer.TransactionCredit {
		return errors.New("NO_CASH")
	} else {
		s.Credit_balance = s.Credit_balance - formTransactionCustomer.TransactionCredit
	}
	return err
}

//預金から残高を引き落とす
