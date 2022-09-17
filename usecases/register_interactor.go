package usecases

import (
	"go_bank/entity"
	"log"
)

const customer_id_length = 16

type setsCustomer entity.Customer

func (f *FormCusotmoer) Register() error {
	var err error

	customerID, err := GenerateUuid(customer_id_length)
	if err != nil {
		//error.Newnに書き換えてreturnするように変更する
		log.Fatal(err)
	}
	accountNumberUuid := GenerateAccountId(f.BranchNumer)
	if err != nil {
		//error.Newnに書き換えてreturnするように変更する

		log.Fatal(err)
	}
	accountNumber := f.BranchNumer + accountNumberUuid[:6]

	Customer := NewCustomer(customerID[:15], accountNumber, f.BranchNumer, f.Name)

	err = Customer.RegisterCustomer()
	if err != nil {
		return err
	}
	return err
}

func NewCustomer(customerID, accountNumber, branchNumber, name string) *Customer {

	return &Customer{
		Customer_id:    customerID,
		Account_number: accountNumber,
		Branch_number:  branchNumber,
		Name:           name,
		Credit_balance: 0,
	}
}

type RegisterInterface interface {
	Register() error
}

type tet struct {
	test string
	tes  int
}
