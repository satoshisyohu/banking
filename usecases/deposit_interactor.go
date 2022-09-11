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
			var updateCreditBalance = strconv.Itoa(caliculateDepositCredit(selectCustomer.Credit_balance, f.TransactionCredit))

			err = selectCustomer.CustomerUpdate(updateCreditBalance)
			if err != nil {
				return errors.New("UPDATE_FAIL")
			}
			creditHistoryInterface := newcreditHistory(f.CustomerId, f.TransactionCredit)
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

func newcreditHistory(CustomerID, TransactionCredit string) CreditHistoryInterface {

	return &CreditHistory{
		Customer_id:        CustomerID,
		Credit_id:          GenerateCreditId(),
		Transaction_credit: TransactionCredit,
		Credit_flag:        "0",
	}
}

type DepositInterface interface {
	Deposit() error
}
