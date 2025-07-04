package repositories

import "api/internal/domain/entities"

type CustomerRepo struct {
	customers []*entities.Customer
	profiles  []*entities.Profile
}

var _ ICustomerRepo = (*CustomerRepo)(nil)

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{
		customers: make([]*entities.Customer, 0),
		profiles:  make([]*entities.Profile, 0),
	}
}

func (r *CustomerRepo) Save(customer *entities.Customer) {
	r.customers = append(r.customers, customer)
}

func (r *CustomerRepo) GetByID(id string) *entities.Customer {
	for _, customer := range r.customers {
		if customer.ID == id {
			return customer
		}
	}

	return nil
}

func (r *CustomerRepo) GetByEmail(email string) *entities.Customer {
	for _, customer := range r.customers {
		if customer.Email == email {
			return customer
		}
	}

	return nil
}

func (r *CustomerRepo) Delete(id string) {
	for i, customer := range r.customers {
		if customer.ID == id {
			r.customers = append(r.customers[:i], r.customers[i+1:]...)
		}
	}
}
