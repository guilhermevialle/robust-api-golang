package entities

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Customer struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=3,max=80,nameregex"`
	Email string `json:"email" validate:"required,email"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("nameregex", func(fl validator.FieldLevel) bool {
		name := fl.Field().String()
		regex := regexp.MustCompile(`^[A-Za-zÀ-ÿ\s]+$`)
		return regex.MatchString(name)
	})
}

func NewCustomer(name, email string) (*Customer, error) {
	id, err := gonanoid.New(21)
	if err != nil {
		return nil, err
	}

	customer := &Customer{
		ID:    id,
		Name:  name,
		Email: email,
	}

	if err := validate.Struct(customer); err != nil {
		return nil, err
	}

	return customer, nil
}
