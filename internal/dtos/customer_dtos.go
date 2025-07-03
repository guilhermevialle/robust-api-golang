package dtos

type CreateCustomerDto struct {
	Name  string `json:"name" validate:"required,min=3,max=80,nameregex"`
	Email string `json:"email" validate:"required,email"`
}
