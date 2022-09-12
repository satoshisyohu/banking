package usecases

import (
	"errors"
	"log"
	"strconv"
)

type TransferInterface interface {
	Transfer(customerId string) error
}

func (t FormTransferCustomer) Transfer(customerId string) error {
	var err error

	fromCustomer, err := IsCustomer(customerId)
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
	// fromCustomer.CustomerUpdate()

	//出金先の相手の情報があるかチェック
	//残高と顧客ID諸々を取得する
	//入出金額の確認をする

	//振り込む側からお金を引き落とす
	//振り込まれた側に金額を追加する
	return err
}
func (c *Customer) isValidTransferCredit(transferCredit int, flag string) error {
	var err error
	fromCreditBalance, err := strconv.Atoi(c.Credit_balance)
	if err != nil {
		return errors.New("INVALID_VALUE")
	}
	switch flag {
	case "1":
		if transferCredit >= fromCreditBalance {
			return errors.New("NO_CASH")
		} else {
			resCreditBalance := fromCreditBalance - transferCredit
			c.Credit_balance = strconv.Itoa(resCreditBalance)
		}
	case "0":
		resCreditBalance := fromCreditBalance + transferCredit
		c.Credit_balance = strconv.Itoa(resCreditBalance)
	}

	return err
}
