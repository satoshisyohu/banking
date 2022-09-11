package usecases

import (
	"go_bank/entity"
	"log"
)

func IsCustomer(CustomerId string) (*entity.Customer, *error) {

	customer := entity.Customer{}

	err := DB.Get(&customer, `SELECT * from customer where customer_id =?`, &CustomerId)
	log.Println(customer)
	if err != nil {
		return &customer, &err
	}
	return &customer, &err
}
