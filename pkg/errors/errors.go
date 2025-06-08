package errors

type ErrorType string

const (
	DomainError     ErrorType = "DOMAIN_ERROR"
	ValidationError ErrorType = "VALIDATION_ERROR"
	NotFoundError   ErrorType = "NOT_FOUND"
	ConflictError   ErrorType = "CONFLICT"
)

type AppError struct {
	Type    ErrorType
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func NewDomainError(code, message string) *AppError {
	return &AppError{
		Type:    DomainError,
		Code:    code,
		Message: message,
	}
}

func NewValidationError(code, message string) *AppError {
	return &AppError{
		Type:    ValidationError,
		Code:    code,
		Message: message,
	}
}

func NewNotFoundError(code, message string) *AppError {
	return &AppError{
		Type:    NotFoundError,
		Code:    code,
		Message: message,
	}
}

func NewConflictError(code, message string) *AppError {
	return &AppError{
		Type:    ConflictError,
		Code:    code,
		Message: message,
	}
}
