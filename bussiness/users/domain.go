package users

import "time"

type Domain struct {
	ID        int
	Username  string
	Password  string
	Name      string
	DOB       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(user *Domain) (*Domain, error)
	Login(username string, password string) (string, error)
	Edit(user *Domain, id int) (*Domain, error)
}

type Repository interface {
	Create(user *Domain) (*Domain, error)
	Update(user *Domain, id int) (*Domain, error)
	FindByUsername(username string) (*Domain, error)
}
