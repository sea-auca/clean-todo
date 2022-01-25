package user

type Repository interface {
	Create(us *User) error
	Update(us *User) error
	Delete(us *User) error

	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
}
