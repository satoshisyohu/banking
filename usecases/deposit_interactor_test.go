package usecases

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCaliculateDepositCredit(t *testing.T) {
	customerId := "dcApU0AwZt_sLa6"
	c := Customer{Branch_number: "02", Account_number: "01015184", Name: "tanaka keisuke", Customer_id: customerId, Credit_balance: 1000000}
	c.CaliculateDepositCredit(10000)
	if c.Credit_balance != 1010000 {
		t.Errorf("expect 1010000,but %d", c.Credit_balance)
	}
}
func TestCaliculateDepositCreditBentchMark(t *testing.B) {
	customerId := "dcApU0AwZt_sLa6"
	c := Customer{Branch_number: "02", Account_number: "01015184", Name: "tanaka keisuke", Customer_id: customerId, Credit_balance: 1000000}
	c.CaliculateDepositCredit(10000)
	if c.Credit_balance != 1010000 {
		t.Errorf("expect 1010000,but %d", c.Credit_balance)
	}
}

func TestDeposit(t *testing.T) {

}
