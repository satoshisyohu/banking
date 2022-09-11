package usecases

func Inquiry(customerId string) (string, *error) {
	customer, err := IsCustomer(customerId)

	if *err != nil {
		return "", err
	}
	return customer.Credit_balance, err
}
