package usecases

import (
	"errors"
	"log"
)

type TransferInterface interface {
	Transfer(FormInquieryCustomer) error
}

func (t FormTransferCustomer) Transfer(formCustomerId FormInquieryCustomer) error {
	var err error
	// var is IsChesckCustomer
	var is IsChesckCustomer
	is = formCustomerId

	// is = formCustomerId
	fromCustomer, err := is.IsCustomer()

	if err != nil {
		return errors.New("NO_CUSTOMER_ID")
	}
	var customerInterface IsCustomerAndCredit
	customerInterface = t
	toCustomer, err := customerInterface.IsCustomerAndCredit()
	log.Println("関数内で更新した現在残高が反映されているか確認")

	log.Println(toCustomer)

	err = fromCustomer.isValidTransferCredit(t.TransferCredit, "1")
	if err != nil {
		return err
	}
	err = fromCustomer.CustomerUpdate()
	if err != nil {
		return err
	}

	err = toCustomer.isValidTransferCredit(t.TransferCredit, "0")
	if err != nil {
		return err
	}
	err = toCustomer.CustomerUpdate()
	if err != nil {
		return err
	}
	log.Println(fromCustomer.Credit_balance)

	return err
}
func (c *Customer) isValidTransferCredit(transferCredit int, flag string) error {
	var err error

	switch flag {
	case "1":
		if transferCredit >= c.Credit_balance {
			return errors.New("NO_CASH")
		} else {
			c.Credit_balance = c.Credit_balance - transferCredit
		}
	case "0":
		c.Credit_balance = c.Credit_balance + transferCredit
	}

	return err
}
