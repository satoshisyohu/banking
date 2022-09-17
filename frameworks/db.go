package frameworks

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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
    customer_id char(16) NOT NULL ,
    credit_id char(8) NOT NULL PRIMARY KEY,
    transaction_credit char(20),
	credit_flag char(1),
	transaction_day timestamp
);`
)

func Config(f *string) {

	if err := godotenv.Load("/Users/Satoshi/Desktop/go_bank/.env"); err != nil {
		log.Println(err)
	}

	connect := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))

	db, err := sqlx.Connect(os.Getenv("DBMS"), connect)
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
