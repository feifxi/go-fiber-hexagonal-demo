package user

type Service interface {
	Register(user *User) error
	GetAll() ([]User, error)
	GetById(id uint) (*User, error)
}