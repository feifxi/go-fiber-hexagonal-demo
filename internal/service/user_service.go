package service

import (
	userDomain "chanombude/super-hexagonal/internal/domain/user"
)

type userService struct {
	repo userDomain.Repository
}

func NewUserService(r userDomain.Repository) userDomain.Service {
	return &userService{repo: r}
}

func (s *userService) Register(user *userDomain.User) error {
	isExists, err := s.repo.ExistsByEmail(user.Email);
	if  err != nil {
		return err
	}
	if isExists {
		return userDomain.ErrEmailAlreadyExists
	}
	return s.repo.Save(user)
}

func (s *userService) GetAll() ([]userDomain.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetById(id uint) (*userDomain.User, error) {
	return s.repo.FindById(id)
}
