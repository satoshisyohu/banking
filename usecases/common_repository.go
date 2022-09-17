package usecases

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB
var Err error

func init() {
	if err := godotenv.Load("/Users/Satoshi/Desktop/go_bank/.env"); err != nil {
		log.Println(err)
	}
	connect := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
	DB, Err = sqlx.Connect(os.Getenv("DBMS"), connect)
	if Err != nil {
		panic(Err)
	}
}

func (c Customer) CustomerUpdate() error {
	tx := DB.MustBegin()
	tx.MustExec(`UPDATE customer SET credit_balance = ? where customer_id = ?`, c.Credit_balance, c.Customer_id)
	defer tx.Rollback()

	return tx.Commit()

}

//この辺はポインタ関数使えるからinterfacaceにできるはず

func (f FormTransactionCreditCustomer) IsCustomerAndCredit() (*Customer, error) {
	customer := Customer{}
	log.Println(f.CustomerId)
	err := DB.Get(&customer, `SELECT * from customer where customer_id =?`, f.CustomerId)
	if err != nil {
		log.Println(customer)
		log.Println(err)

	}

	return &customer, err
}

func (t FormTransferCustomer) IsCustomerAndCredit() (*Customer, error) {
	customer := Customer{}
	err := DB.Get(&customer, `SELECT * from customer where branch_number = ? and account_number=?`, t.BranchNumer, t.AccountNumber)
	log.Println(customer)

	return &customer, err
}

func (t FormInquieryCustomer) IsCustomer() (*Customer, error) {
	customer := Customer{}
	err := DB.Get(&customer, `SELECT * from customer where customer_id =?`, t.CustomerId)
	log.Println(customer)

	return &customer, err
}

func (c *CreditHistory) RegisterTransacationHistory() error {

	tx := DB.MustBegin()

	// tx.MustExec("INSERT INTO customer(customer_id,account_number,branch_number,name,credit_balance) VALUES($1,$2,$3,$4,$5)", customer.Customer_id, customer.Account_number, customer.Branch_number, customer.Name, customer.Credit_balance)
	tx.MustExec(`INSERT INTO credit_history(customer_id,credit_id,transaction_credit,credit_flag,transaction_day) VALUES(?,?,?,?,NOW())`, c.Customer_id, c.Credit_id, c.Transaction_credit, c.Credit_flag)
	err := tx.Commit()
	if err != nil {
		return err
	} else {
		log.Print("succed credit_history customer")
	}
	defer tx.Rollback()

	return err
}

type CreditHistoryInterface interface {
	RegisterTransacationHistory() error
}

type IsCustomerAndCredit interface {
	IsCustomerAndCredit() (*Customer, error)
}
type IsCheckCustomer interface {
	IsCustomer() (*Customer, error)
}
