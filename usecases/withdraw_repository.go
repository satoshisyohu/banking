package usecases

import (
	"go_bank/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

// type WithDraw interface {
// 	IsCustomerAndCredit(credit_history_form *entity.Credit_history) bool
// 	RegisterTransacationHistory(credit_history_form *entity.Credit_history)
// }

type TestC entity.Customer

func (c *testC) Test() bool {
	return c.Account_number == "aa"
}

func IsCustomerAndCredit(creditHistoryForm *entity.Credit_history) *entity.Customer {
	db, err := ConnectToDb()
	customer := entity.Customer{}
	if err != nil {
		log.Fatal(err)
	} else {
		db.Get(&customer, `SELECT * from customer where customer_id =?`, creditHistoryForm.Customer_id)
	}

	return &customer
}

func RegisterTransacationHistory(set_credit_history *entity.Credit_history) {
	db, err := ConnectToDb()
	if err != nil {
		log.Print(err)
	} else {

		tx := db.MustBegin()

		// tx.MustExec("INSERT INTO customer(customer_id,account_number,branch_number,name,credit_balance) VALUES($1,$2,$3,$4,$5)", customer.Customer_id, customer.Account_number, customer.Branch_number, customer.Name, customer.Credit_balance)
		tx.MustExec(`INSERT INTO credit_history(customer_id,credit_id,transaction_credit,credit_flag,transaction_day) VALUES(?,?,?,?,CURDATE())`, set_credit_history.Customer_id, set_credit_history.Credit_id, set_credit_history.Transaction_credit, set_credit_history.Credit_flag)
		tx.Commit()
		log.Print("succed credit_history customer")
	}

	//todo customerの残高も更新すること
}

func CustomerUpdate(customerId, creditBalance string) error {
	db, err := ConnectToDb()
	if err != nil {
		return err
	} else {

		tx := db.MustBegin()
		tx.MustExec(`UPDATE customer SET credit_balance = ? where customer_id = ?`, creditBalance, customerId)
		tx.Commit()
		log.Print("update customer credit_balance")
	}
	return err
}

func ConnectToDb() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", "root:root@tcp(0.0.0.0:3333)/banking_db")
}
