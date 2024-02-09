package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"bank/repository"
)

func main() {
	db, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/bank?parseTime=true")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)

	_ = customerRepository
	customers, err := customerRepository.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
}
