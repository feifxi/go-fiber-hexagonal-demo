package validator

import (
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

	// Check length
	if len(password) < 8 {
		return false
	}

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
