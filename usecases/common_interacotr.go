package usecases

import (
	"encoding/base64"
	"go_bank/entity"
	"math/rand"
	"strconv"
	"time"
)

const CREDIT_ID_LENGTH = 8      //頭にbranch_nameを足すために6桁にしている
const account_number_length = 6 //頭にbranch_nameを足すために6桁にしている

type Customer entity.Customer
type CreditHistory entity.Credit_history

func GenerateCreditId() string {
	var creditId = ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < CREDIT_ID_LENGTH; i++ {
		creditId = creditId + strconv.Itoa(rand.Intn(10))
	}
	return creditId
}

func GenerateUuid(length int) (string, error) {
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func GenerateAccountId(account_number string) string {
	for i := 0; i < account_number_length; i++ {
		account_number = account_number + strconv.Itoa(rand.Intn(10))
	}
	return account_number
}
