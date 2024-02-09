package repository

import (
	"errors"
	"time"
)

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{
			CustomerID:  1,
			Name:        "John",
			DateOfBirth: time.Now(),
			City:        "New York",
			ZipCode:     "12345",
			Status:      "active",
		},
		{
			CustomerID:  2,
			Name:        "Jane",
			DateOfBirth: time.Now(),
			City:        "Los Angeles",
			ZipCode:     "67890",
			Status:      "active",
		},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
