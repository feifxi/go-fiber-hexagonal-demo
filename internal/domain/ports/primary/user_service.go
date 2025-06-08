package primary

import (
	"chanombude/super-hexagonal/internal/domain/model"
)

type UserService interface {
	Register(name, email, password string) error
	GetAll() ([]model.User, error)
	GetById(id uint) (*model.User, error)
}
