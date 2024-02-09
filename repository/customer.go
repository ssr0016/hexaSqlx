package repository

import "time"

type Customer struct {
	CustomerID  int       `db:"customer_id"`
	Name        string    `db:"name"`
	DateOfBirth time.Time `db:"date_of_birth"`
	City        string    `db:"city"`
	ZipCode     string    `db:"zip_code"`
	Status      string    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}
