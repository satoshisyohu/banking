package entity

type Customer struct {
	Customer_id    *string `json:"customer_id"`
	Account_number string  `json:"account_number"`
	Branch_number  string  `json:"branch_number" validate:"required"`
	Name           string  `json:"name" validate:"required"`
	Credit_balance int     `json:"credit_balance"`
}

type Credit_history struct {
	Customer_id        string `json:"customer_id" validate:"required"`
	Credit_id          string `json:"credit_id"`
	Transaction_credit int    `json:"transaction_credit" validate:"required"`
	Credit_flag        string `json:"credit_flag"` //入金:0 出金1
	Transaction_day    int    `json:"transaction_day"`
}

type ResultMessage struct {
	Result      bool
	MessageType Result
}
type Result int

const (
	WITHDRAW_OK Result = iota
	NO_CUSTOMER_ID
	NO_CASH
	INVALID_VALUE
	UPDATE_FAIL
	DEPOSIT_OK
)

type ReturnResult struct {
	ResultMessage string `json:"result_message"`
}

type ReturnCredit struct {
	ResultCrecit int `json:"result_credit"`
}
