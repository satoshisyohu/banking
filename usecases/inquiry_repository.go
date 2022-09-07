package usecases

import (
	"go_bank/entity"
	"log"
)

func IsCustomer(CustomerId string) *entity.Customer {
	db, err := ConnectToDb()
	customer := entity.Customer{}
	log.Println(&CustomerId)
	log.Println(CustomerId)

	if err != nil {
		log.Panic(err)
	} else {
		db.Get(&customer, `SELECT * from customer where customer_id =?`, &CustomerId)
	}
	log.Println(customer)
	return &customer
}
