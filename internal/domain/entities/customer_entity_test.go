package entities_test

import (
	"api/internal/domain/entities"
	"testing"
)

func TestCustomer(t *testing.T) {
	t.Run("should create a new valid customer", func(t *testing.T) {
		if _, err := entities.NewCustomer("John Doe", "j@j.com"); err != nil {
			t.Error(err)
		}
	})

	t.Run("should not create a new invalid customer", func(t *testing.T) {
		if _, err := entities.NewCustomer("", "j@j.com"); err == nil {
			t.Error("expected error")
		}

		if _, err := entities.NewCustomer("John Doe", ""); err == nil {
			t.Error("expected error")
		}

		if _, err := entities.NewCustomer("2ewqds313", "j"); err == nil {
			t.Error("expected error")
		}

		if _, err := entities.NewCustomer("2", "j@j."); err == nil {
			t.Error("expected error")
		}

		if _, err := entities.NewCustomer("John Doe", "j@j"); err == nil {
			t.Error("expected error")
		}

	})
}
