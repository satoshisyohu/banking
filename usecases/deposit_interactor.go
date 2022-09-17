package usecases

import (
	"errors"
	"log"
)

// 預金から残高を引き落とす
func (f FormTransactionCreditCustomer) Deposit() error {
	var err error

	selectCustomer, err := f.IsCustomerAndCredit()
	log.Println(selectCustomer)
	if err != nil {
		return errors.New("NO_CUSTOMER_ID")
	} else {
		selectCustomer.CaliculateDepositCredit(f.TransactionCredit)

		err = selectCustomer.CustomerUpdate()
		if err != nil {
			return errors.New("UPDATE_FAIL")
		}
		err := f.NewCreditHistory().RegisterTransacationHistory()
		if err != nil {
			return err
		}
	}
	return err
}

//預金から残高を引き落とす
func (c *Customer) CaliculateDepositCredit(formTransactionCredit int) {
	c.Credit_balance = c.Credit_balance + formTransactionCredit
}

type DepositInterface interface {
	Deposit() error
}
