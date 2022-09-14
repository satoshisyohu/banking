package usecases

import (
	"errors"
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

		selectCustomer.caliculateDepositCredit(f.TransactionCredit)

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

//預金から残高を引き落とす
func (c *Customer) caliculateDepositCredit(formTransactionCredit int) {
	c.Credit_balance = c.Credit_balance + formTransactionCredit
}

type DepositInterface interface {
	Deposit() error
}
