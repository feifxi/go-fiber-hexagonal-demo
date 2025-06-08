package dto

import (
	"reflect"
)

type (
	RegisterUserRequest struct {
		Name     string `json:"name" validate:"required,name" vmsg:"Name must be between 2-50 characters and contain only letters, spaces, and hyphens"`
		Email    string `json:"email" validate:"required,email" vmsg:"Please provide a valid email address"`
		Password string `json:"password" validate:"required,password" vmsg:"Password must be at least 8 characters long and contain uppercase, lowercase, number, and special character"`
	}

	UserResponse struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email" vmsg:"Please provide a valid email address"`
		Password string `json:"password" validate:"required" vmsg:"Password is required"`
	}

	UpdateUserRequest struct {
		Name     string `json:"name" validate:"omitempty,name" vmsg:"Name must be between 2-50 characters and contain only letters, spaces, and hyphens"`
		Email    string `json:"email" validate:"omitempty,email" vmsg:"Please provide a valid email address"`
		Password string `json:"password" validate:"omitempty,password" vmsg:"Password must be at least 8 characters long and contain uppercase, lowercase, number, and special character"`
	}
)

// GetValidationErrors returns a map of field names to error messages
func (r *RegisterUserRequest) GetValidationErrors() map[string]string {
	errors := make(map[string]string)

	// Get struct type
	t := reflect.TypeOf(*r)

	// Iterate through fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get validation message
		if msg, ok := field.Tag.Lookup("vmsg"); ok {
			errors[field.Name] = msg
		}
	}

	return errors
}
