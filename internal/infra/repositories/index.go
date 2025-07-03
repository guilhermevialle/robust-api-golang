package repositories

import "api/internal/domain/entities"

type ICustomerRepo interface {
	Save(customer *entities.Customer)
	GetByID(id string) *entities.Customer
	GetByEmail(email string) *entities.Customer
	Delete(id string)
}
