package response

import (
	"time"
	"todolist/bussiness/users"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	DOB       string    `json:"dob"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Token string `json:"token"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:        domain.ID,
		Username:  domain.Username,
		Password:  domain.Password,
		Name:      domain.Name,
		DOB:       domain.DOB,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
