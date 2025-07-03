package repositories

import "api/internal/domain/entities"

type CustomerRepo struct {
	storage []*entities.Customer
}

var _ ICustomerRepo = (*CustomerRepo)(nil)

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{
		storage: make([]*entities.Customer, 0),
	}
}

func (r *CustomerRepo) Save(customer *entities.Customer) {
	r.storage = append(r.storage, customer)
}

func (r *CustomerRepo) GetByID(id string) *entities.Customer {
	for _, customer := range r.storage {
		if customer.ID == id {
			return customer
		}
	}

	return nil
}

func (r *CustomerRepo) GetByEmail(email string) *entities.Customer {
	for _, customer := range r.storage {
		if customer.Email == email {
			return customer
		}
	}

	return nil
}

func (r *CustomerRepo) Delete(id string) {
	for i, customer := range r.storage {
		if customer.ID == id {
			r.storage = append(r.storage[:i], r.storage[i+1:]...)
		}
	}
}
