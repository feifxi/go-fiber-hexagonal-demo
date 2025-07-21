package port

import "chanombude/super-hexagonal/internal/model"

type UserService interface {
	Register(user *model.User) error
	GetAll() ([]model.User, error)
	GetById(id uint) (*model.User, error)
}

type UserRepository interface {
	Save(user *model.User) error
	FindAll() ([]model.User, error)
	FindById(id uint) (*model.User, error)
	ExistsByEmail(email string) (bool, error)
}