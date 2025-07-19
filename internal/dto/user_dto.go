package dto

type (
	RegisterUserRequest struct {
		Name     string `json:"name" validate:"required,min=3,name" vmsg:"Name must be between 2-50 characters and contain only letters, spaces, and hyphens"`
		Email    string `json:"email" validate:"required,email" vmsg:"Please provide a valid email address"`
		Password string `json:"password" validate:"required,min=8,max=30,password" vmsg:"Password must be contain uppercase, lowercase, number, and special character"`
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

