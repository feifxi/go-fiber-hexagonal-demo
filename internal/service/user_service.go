package service

import (
	"chanombude/super-hexagonal/internal/domain/errors"
	"chanombude/super-hexagonal/internal/domain/model"
	"chanombude/super-hexagonal/internal/domain/ports/primary"
	"chanombude/super-hexagonal/internal/domain/ports/secondary"
)

type userService struct {
	userRepo secondary.UserRepository
}

func NewUserService(repo secondary.UserRepository) primary.UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) Register(name, email, password string) error {
	exists, err := s.userRepo.ExistsByEmail(email)
	if err != nil {
		return err
	}
	if exists {
		return errors.ErrEmailAlreadyExists
	}

	user, err := model.NewUser(name, email, password)
	if err != nil {
		return err
	}

	return s.userRepo.Save(user)
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetById(id uint) (*model.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrUserNotFound
	}
	return user, nil
}
