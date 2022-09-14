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

			// var updateCreditBalance = strconv.Itoa(caliculateWithdrawCredit(f.TransactionCredit))
			err = selectCustomer.caliculateWithdrawCredit(f.TransactionCredit)
			if err != nil {
				return err
			}
			err = selectCustomer.CustomerUpdate()
			if err != nil {
				return errors.New("UPDATE_FAIL")

			}
			creditHistoryInterface := f.NewCreditHistory()
			err := creditHistoryInterface.RegisterTransacationHistory()
			if *err != nil {
				log.Println("test")

				log.Println("test")

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
func (c *Customer) caliculateWithdrawCredit(formTransactionCredit string) error {
	var err error
	intSelectCredit, _ := strconv.Atoi(c.Credit_balance)
	intFormCreditBalance, _ := strconv.Atoi(formTransactionCredit)
	c.Credit_balance = strconv.Itoa(intSelectCredit - intFormCreditBalance)
	if err != nil {
		return err
	}
	return err
}

type WithdrawInterface interface {
	Withdraw() error
}
