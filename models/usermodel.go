package models

type User struct {
	Email string `db:"email"`
}

type UserStore interface {
	User(Email string) (User, error)
	Users() ([]User, error)
	CreateUser(u *User) error
}
