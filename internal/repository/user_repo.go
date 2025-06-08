package repository

import (
	"chanombude/super-hexagonal/internal/domain/user"
	"errors"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepo{db: db}
}

func (r *userRepo) Save(u *user.User) error {
	return r.db.Create(u).Error
}

func (r *userRepo) FindAll() ([]user.User, error) {
	var users []user.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) FindById(id uint) (*user.User, error) {
	var user user.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err // some other error
	}
	return &user, nil
}

func (r *userRepo) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&user.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}