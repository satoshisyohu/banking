package entity

type FormCustomer struct {
	BranchNumer string `json:"branchNumber" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type FormInquieryCustomer struct {
	CustomerId string `json:"customerId" validate:"required"`
}

//顧客の取引を行う際に受け取るための構造体
type FormTransactionCustomer struct {
	CustomerId        string `json:"customerId" validate:"required"`
	TransactionCredit int    `json:"transactionCredit" validate:"required"`
}

// TransactionValidation implements controllers.ValidationTransactionInterface

//支店番号
//講座番号
//金額
type FormTransferCustomer struct {
	BranchNumer    string `json:"branchNumber" validate:"required"`
	AccountNumber  string `json:"accountNumber" validate:"required"`
	TransferCredit int    `json:"transferCredit" validate:"required"`
}

type test struct {
	branch string
}
