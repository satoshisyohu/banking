package usecases

func (c FormInquieryCustomer) Inquiry() (int, error) {

	var iscustomer IsCheckCustomer = c
	customer, err := iscustomer.IsCustomer()

	if err != nil {
		return 0, err
	}
	return customer.Credit_balance, err
}

type InquiryInterface interface {
	Inquiry() (int, error)
}
