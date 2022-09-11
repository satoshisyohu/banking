package usecases

import (
	"log"
)

func (c Customer) RegisterCustomer() error {
	var err error

	tx := DB.MustBegin()
	//MustExecは戻り値にerrを指定すると必ずerrに値がなにか入るので、うまく処理できない
	//errorが起こった際はpanicになるから問題なし
	//でも本番でpanicが起こったらサーバ止まるよな
	tx.MustExec(`INSERT INTO customer(customer_id,account_number,branch_number,name,credit_balance) VALUES(?,?,?,?,?)`, c.Customer_id, c.Account_number, c.Branch_number, c.Name, c.Credit_balance)
	err = tx.Commit()
	if err != nil {
		return err
	} else {
		log.Print("succeed register customer")
	}
	defer tx.Rollback()

	return err
}
