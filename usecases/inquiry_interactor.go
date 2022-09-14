package usecases

func (c FormInquieryCustomer) Inquiry() (string, error) {
	var iscustomer IsChesckCustomer
	iscustomer = c
	customer, err := iscustomer.IsCustomer()

	if err != nil {
		return "", err
	}
	return customer.Credit_balance, err
}

type InquiryInterface interface {
	Inquiry() (string, error)
}
