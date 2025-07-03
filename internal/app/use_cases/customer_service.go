package use_cases

import (
	"api/internal/domain/entities"
	"api/internal/infra/repositories"
	"errors"
)

type ICustomerService interface {
	CreateCustomer(name, email string) (*entities.Customer, error)
	DeleteCustomer(id string) error
	GetCustomerByEmail(email string) (*entities.Customer, error)
}

type CustomerService struct {
	customerRepo repositories.ICustomerRepo
}

var _ ICustomerService = (*CustomerService)(nil)

func NewCustomerService(customerRepo repositories.ICustomerRepo) *CustomerService {
	return &CustomerService{
		customerRepo: customerRepo,
	}
}

func (cs *CustomerService) CreateCustomer(name, email string) (*entities.Customer, error) {
	if cs.customerRepo.GetByEmail(email) != nil {
		return nil, errors.New("customer with email already exists")
	}

	customer, err := entities.NewCustomer(name, email)
	if err != nil {
		return nil, errors.New("Bad request")
	}

	cs.customerRepo.Save(customer)

	return customer, nil
}

func (cs *CustomerService) DeleteCustomer(id string) error {
	if cs.customerRepo.GetByID(id) == nil {
		return errors.New("customer not found")
	}

	cs.customerRepo.Delete(id)

	return nil
}

func (cs *CustomerService) GetCustomerByEmail(email string) (*entities.Customer, error) {
	customer := cs.customerRepo.GetByEmail(email)
	if customer == nil {
		return nil, errors.New("customer not found")
	}

	return customer, nil
}
