package secondary

import (
	"chanombude/super-hexagonal/internal/domain/model"
)

type UserRepository interface {
	Save(user *model.User) error
	FindAll() ([]model.User, error)
	FindById(id uint) (*model.User, error)
	ExistsByEmail(email string) (bool, error)
}
