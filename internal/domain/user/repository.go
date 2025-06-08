package user

type Repository interface {
	Save(user *User) error
	FindAll() ([]User, error)
	FindById(id uint) (*User, error)
	ExistsByEmail(email string) (bool, error)
}