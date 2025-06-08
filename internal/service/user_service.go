package service

import (
	"chanombude/super-hexagonal/internal/domain"
	"chanombude/super-hexagonal/internal/pkg/errors"
	"chanombude/super-hexagonal/internal/repository"
)

type UserService interface {
	Register(user *domain.User) error
	GetAll() ([]domain.User, error)
	GetById(id uint) (*domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) Register(user *domain.User) error {
	exists, err := s.userRepo.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.NewConflictError("CONFLICT_EMAIL", "this email already registered")
	}

	newUser, err := domain.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		return errors.NewValidationError("INVALID_PASSWORD", "invalid password format")
	}

	if err := s.userRepo.Save(newUser); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetAll() ([]domain.User, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetById(id uint) (*domain.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
