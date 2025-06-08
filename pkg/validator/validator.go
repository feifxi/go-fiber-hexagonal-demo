package validator

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func Init() {
	validate = validator.New()

	// Register custom validations
	registerCustomValidations(validate)
}

func Get() *validator.Validate {
	return validate
}

func ValidateStruct(s interface{}) map[string]string {	
	err := validate.Struct(s)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// Create a map to store field-specific error messages
		errorsMap := make(map[string]string)
		for _, fieldError := range validationErrors {
			// Customize the error message based on the field and tag
			switch fieldError.Tag() {
			case "required":
				errorsMap[fieldError.Field()] = fmt.Sprintf("%s is required.", fieldError.Field())
			case "min":
				errorsMap[fieldError.Field()] = fmt.Sprintf("%s must be at least %s characters long.", fieldError.Field(), fieldError.Param())
			case "max":
				errorsMap[fieldError.Field()] = fmt.Sprintf("%s cannot exceed %s characters.", fieldError.Field(), fieldError.Param())
			case "email":
				errorsMap[fieldError.Field()] = fmt.Sprintf("%s must be a valid email address.", fieldError.Field())
			case "alpha":
				errorsMap[fieldError.Field()] = fmt.Sprintf("%s must contain only alphabetic characters.", fieldError.Field())
			// Add more cases for other validation tags as needed
			default:
				errorsMap[fieldError.Field()] = getCustomErrorMessage(s, fieldError.Field())
			}
		}
		return errorsMap
	}
	return nil
}

func getCustomErrorMessage(s interface{}, field string) string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if f, ok := t.FieldByName(field); ok {
		if msg, ok := f.Tag.Lookup("vmsg"); ok {
			return msg
		}
	}
	return fmt.Sprintf("invalid %s field.", field)
}

func registerCustomValidations(v *validator.Validate) {
	// Register password validation
	v.RegisterValidation("password", validatePassword)

	// Register name validation
	v.RegisterValidation("name", validateName)
}

// validatePassword checks if password meets requirements:
// - At least 8 characters
// - At least one uppercase letter
// - At least one lowercase letter
// - At least one number
// - At least one special character
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check for uppercase
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpper {
		return false
	}

	// Check for lowercase
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !hasLower {
		return false
	}

	// Check for number
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		return false
	}

	// Check for special character
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)
	if !hasSpecial {
		return false
	}

	return true
}

// validateName checks if name meets requirements:
// - Only letters, spaces, and hyphens
// - Between 2 and 50 characters
func validateName(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	// Check length
	if len(name) < 2 || len(name) > 50 {
		return false
	}

	// Check for valid characters
	validName := regexp.MustCompile(`^[a-zA-Z\s-]+$`).MatchString(name)
	if !validName {
		return false
	}

	return true
}
