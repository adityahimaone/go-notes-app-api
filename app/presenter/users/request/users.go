package request

import (
	"todolist/bussiness/users"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	DOB      string `json:"dob"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToDomain(request User) *users.Domain {
	return &users.Domain{
		Username: request.Username,
		Password: request.Password,
		Name:     request.Name,
		DOB:      request.DOB,
	}
}
