package entities_test

import (
	"api/internal/domain/entities"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	t.Run("should create a valid customer", func(t *testing.T) {
		name := "John Doe"
		email := "john@example.com"
		password := "securepassword123"

		customer, err := entities.NewCustomer(name, email, password)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if customer.ID == "" {
			t.Error("expected generated ID, got empty string")
		}

		if customer.Name != name {
			t.Errorf("expected name %s, got %s", name, customer.Name)
		}

		if customer.Email != email {
			t.Errorf("expected email %s, got %s", email, customer.Email)
		}

		if customer.Password != password {
			t.Errorf("expected password to be set correctly")
		}
	})
}
