package service

import (
	"chanombude/super-hexagonal/internal/model"
	"chanombude/super-hexagonal/internal/port"
	"chanombude/super-hexagonal/pkg/errors"
)

type userService struct {
	userRepo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) Register(user *model.User) error {
	exists, err := s.userRepo.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.NewConflictError("CONFLICT_EMAIL", "this email already registered")
	}

	newUser, err := model.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		return errors.NewValidationError("INVALID_PASSWORD", "invalid password format")
	}

	if err := s.userRepo.Save(newUser); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetAll() ([]model.User, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetById(id uint) (*model.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
