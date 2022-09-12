package usecases

import (
	"errors"

	"strconv"
)

// 預金から残高を引き落とす
func (f FormTransactionCreditCustomer) Deposit() error {
	var err error

	//container/listっていうライブラリを使うこともできるよ
	// l := list.New()
	// l.PushBack("z")
	// log.Println(l.Front().Value)

	selectCustomer, err := f.IsCustomerAndCredit()
	if err != nil {
		return errors.New("NO_CUSTOMER_ID")
	} else {
		err = selectCustomer.isValidDepositCredit(f)
		if err != nil {
			return err
		} else {
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
			if err != nil {
				return *err
			}
		}
		return err
	}
}
func (c *Customer) isValidDepositCredit(formTransactionCustomer FormTransactionCreditCustomer) error {
	_, err := strconv.Atoi(formTransactionCustomer.TransactionCredit)
	if err != nil {
		return errors.New("INVALID_VALUE")
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

type DepositInterface interface {
	Deposit() error
}
