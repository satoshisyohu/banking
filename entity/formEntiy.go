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
	TransactionCredit string `json:"transactionCredit" validate:"required"`
}

// TransactionValidation implements controllers.ValidationTransactionInterface
