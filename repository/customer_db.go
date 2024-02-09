package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{
		db: db,
	}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customer := []Customer{}
	query := "SELECT customer_id, name, city, zipcode, date_of_birth FROM customers"
	err := r.db.Select(&customer, query)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT customer_id, name, city, zipcode, date_of_birth FROM customers WHERE customer_id = ?"
	err := r.db.Select(&customer, query, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
