package repository

import (
	"chanombude/super-hexagonal/internal/domain"
	"chanombude/super-hexagonal/internal/pkg/errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *domain.User) error
	FindAll() ([]domain.User, error)
	FindById(id uint) (*domain.User, error)
	ExistsByEmail(email string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.NewDomainError("DB_ERROR", err.Error())
	}
	return nil
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, errors.NewDomainError("DB_ERROR", err.Error())
	}
	return users, nil
}

func (r *userRepository) FindById(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("DB_ERROR", "user not found")
		}
		return nil, errors.NewDomainError("DB_ERROR", err.Error())
	}
	return &user, nil
}

func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&domain.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, errors.NewDomainError("DB_ERROR", err.Error())
	}
	return count > 0, nil
}
