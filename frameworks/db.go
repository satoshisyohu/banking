package frameworks

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	customer_table = `
	CREATE TABLE customer (
    customer_id char(16) NOT NULL PRIMARY KEY,
    account_number char(8),
	branch_number char(2),
    name char(20),
	credit_balance char(20)
);`
	credit_history_table = `
	CREATE TABLE credit_history (
    customer_id char(16) NOT NULL PRIMARY KEY,
    credit_id char(8),
    transaction_credit char(20),
	credit_flag char(1),
	transaction_day date
);`
)

func Config(f *string) {

	db, err := sqlx.Connect("mysql", "root:root@tcp(0.0.0.0:3333)/banking_db")
	if err != nil {
		log.Print(err)
	} else {
		log.Print("db connected")
	}

	//go run main.go実行時に -db=initが指定されている場合dbを作成する
	//指定されていない場合は作成しない
	if *f == "init" {
		log.Println("make Customer tables")
		db.MustExec(customer_table)
		log.Println("make Customer tables done")

		log.Println("make credit_history tables")
		db.MustExec(credit_history_table)
		log.Println("make credit_history table done")
	}
}
