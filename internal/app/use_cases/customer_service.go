package use_cases

import (
	"api/internal/domain/entities"
	"api/internal/infra/repositories"
	"errors"
)

type ICustomerService interface {
	CreateCustomer(name, email, password string) (*entities.Customer, error)
	GetCustomerByEmail(email string) (*entities.Customer, error)
	GetProfileByID(id string) (*entities.Profile, error)
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

func (cs *CustomerService) CreateCustomer(name, email, password string) (*entities.Customer, error) {
	if cs.customerRepo.GetByEmail(email) != nil {
		return nil, errors.New("email already in use")
	}

	customer, err := entities.NewCustomer(name, email, password)
	if err != nil {
		return nil, errors.New("bad request")
	}

	cs.customerRepo.Save(customer)

	return customer, nil
}

func (cs *CustomerService) GetCustomerByEmail(email string) (*entities.Customer, error) {
	customer := cs.customerRepo.GetByEmail(email)
	if customer == nil {
		return nil, errors.New("customer not found")
	}

	return customer, nil
}

func (cs *CustomerService) GetProfileByID(id string) (*entities.Profile, error) {
	customer := cs.customerRepo.GetByID(id)

	if customer == nil {
		return nil, errors.New("customer not found")
	}

	return customer.Profile, nil
}
