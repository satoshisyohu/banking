package usecases

import (
	"go_bank/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func RegisterCustomer(customer *entity.Customer) {
	db, err := sqlx.Connect("mysql", "root:root@tcp(0.0.0.0:3333)/banking_db")
	if err != nil {
		log.Fatal(err)
	} else {
		tx := db.MustBegin()
		//MustExecは戻り値にerrを指定すると必ずerrに値がなにか入るので、うまく処理できない
		//errorが起こった際はpanicになるから問題なし
		//でも本番でpanicが起こったらサーバ止まるよな
		tx.MustExec(`INSERT INTO customer(customer_id,account_number,branch_number,name,credit_balance) VALUES(?,?,?,?,?)`, customer.Customer_id, customer.Account_number, customer.Branch_number, customer.Name, customer.Credit_balance)
		tx.Commit()
		log.Print("succed register customer")

	}
}
