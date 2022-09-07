package usecases

import (
	"encoding/base64"
	"go_bank/entity"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const customer_id_length = 16
const account_number_length = 6 //頭にbranch_nameを足すために6桁にしている

type setsCustomer entity.Customer

type RegisterInterface interface {
	Register()
}

func Register(customer_form *entity.Customer) {

	customerID, err := generateUuid(customer_id_length)
	if err != nil {
		log.Print(err)
	}
	accountNumberUuid := generateAccountId(customer_form.Branch_number)
	if err != nil {
		log.Print(err)
	}
	accountNumber := customer_form.Branch_number + accountNumberUuid[:6]

	var customer setsCustomer
	customer = setsCustomer(*customer_form)
	RegisterCustomer(customer.setCustomer(customerID[:15], accountNumber))

}

func generateUuid(length int) (string, error) {
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func generateAccountId(account_number string) string {
	for i := 0; i < account_number_length; i++ {
		account_number = account_number + strconv.Itoa(rand.Intn(10))
	}
	return account_number
}

func (c *setsCustomer) setCustomer(customerID, accountNumber string) *entity.Customer {

	return &entity.Customer{
		Customer_id:    customerID,
		Account_number: accountNumber,
		Branch_number:  c.Branch_number,
		Name:           c.Name,
		Credit_balance: "0",
	}
}
