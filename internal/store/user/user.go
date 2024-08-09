package user

type User struct {
	ID       uint
	Email    string
	Password string
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}
