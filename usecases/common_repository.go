package usecases

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var Err error

func init() {
	DB, Err = sqlx.Connect("mysql", "root:root@tcp(0.0.0.0:3333)/banking_db")
	if Err != nil {
		panic(Err)
	}
}

func (c Customer) CustomerUpdate(creditBalance string) error {
	tx := DB.MustBegin()
	tx.MustExec(`UPDATE customer SET credit_balance = ? where customer_id = ?`, creditBalance, c.Customer_id)
	defer tx.Rollback()

	return tx.Commit()

}

//この辺はポインタ関数使えるからinterfacaceにできるはず

func (f FormTransactionCreditCustomer) IsCustomerAndCredit() (*Customer, error) {
	customer := Customer{}
	err := DB.Get(&customer, `SELECT * from customer where customer_id =?`, f.CustomerId)
	log.Println(customer)

	return &customer, err
}

func (c *CreditHistory) RegisterTransacationHistory() *error {

	tx := DB.MustBegin()

	// tx.MustExec("INSERT INTO customer(customer_id,account_number,branch_number,name,credit_balance) VALUES($1,$2,$3,$4,$5)", customer.Customer_id, customer.Account_number, customer.Branch_number, customer.Name, customer.Credit_balance)
	tx.MustExec(`INSERT INTO credit_history(customer_id,credit_id,transaction_credit,credit_flag,transaction_day) VALUES(?,?,?,?,CURDATE())`, c.Customer_id, c.Credit_id, c.Transaction_credit, c.Credit_flag)
	err := tx.Commit()
	if err != nil {
		return &err
	} else {
		log.Print("succed credit_history customer")
	}
	defer tx.Rollback()

	return &err
}

type CreditHistoryInterface interface {
	RegisterTransacationHistory() *error
}
