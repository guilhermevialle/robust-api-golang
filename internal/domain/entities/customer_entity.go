package entities

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Customer struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Profile  *Profile `json:"profile"`
}

func NewCustomer(name, email, password string) (*Customer, error) {
	id, err := gonanoid.New(21)
	if err != nil {
		return nil, err
	}

	customer := &Customer{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	p, err := customer.createProfile()
	if err != nil {
		return nil, err
	}

	customer.Profile = p

	return customer, nil
}

func (c *Customer) createProfile() (*Profile, error) {
	id, err := gonanoid.New(21)
	if err != nil {
		return nil, err
	}

	return &Profile{
		ID:         id,
		CustomerID: c.ID,
		Name:       c.Name,
		Summary:    "",
	}, nil
}
