package usecases

func (c FormInquieryCustomer) Inquiry() (string, error) {
	customer, err := IsCustomer(c.CustomerId)

	if err != nil {
		return "", err
	}
	return customer.Credit_balance, err
}

type InquiryInterface interface {
	Inquiry() (string, error)
}
