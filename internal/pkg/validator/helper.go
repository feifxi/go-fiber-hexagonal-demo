package validator

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"reflect"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

// ValidateStruct validates a struct and returns validation errors
func ValidateStruct(s interface{}) []ValidationError {
	var errors []ValidationError
	
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			
			// Get custom message if available
			message := getErrorMessage(s, field, tag)
			if message == "" {
				message = getDefaultMessage(field, tag)
			}
			
			errors = append(errors, ValidationError{
				Field:   field,
				Message: message,
			})
		}
	}
	
	return errors
}

// getErrorMessage gets the custom error message from the struct tag
func getErrorMessage(s interface{}, field, tag string) string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	
	if f, ok := t.FieldByName(field); ok {
		if msg, ok := f.Tag.Lookup("vmsg"); ok {
			return msg
		}
	}
	
	return ""
}

// getDefaultMessage returns a default error message
func getDefaultMessage(field, tag string) string {
	field = strings.ToLower(field)
	
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return "invalid email format"
	case "password":
		return "password must be at least 8 characters long and contain uppercase, lowercase, number, and special character"
	case "name":
		return "name must be between 2-50 characters and contain only letters, spaces, and hyphens"
	default:
		return field + " is invalid"
	}
} 