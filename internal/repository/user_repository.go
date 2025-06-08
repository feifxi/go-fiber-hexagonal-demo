package repository

import (
	"chanombude/super-hexagonal/internal/domain/model"
	"chanombude/super-hexagonal/internal/domain/ports/secondary"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) secondary.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
