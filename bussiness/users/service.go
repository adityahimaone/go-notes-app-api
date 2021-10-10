package users

import (
	"todolist/app/middleware/auth"
	bcrypthelper "todolist/helper/bcrypt"
)

type serviceUsers struct {
	repository Repository
	jwtAuth    *auth.ConfigJWT
}

func NewService(repoUser Repository, jwtauth *auth.ConfigJWT) Service {
	return &serviceUsers{
		repository: repoUser,
		jwtAuth:    jwtauth,
	}
}

func (service serviceUsers) Register(user *Domain) (*Domain, error) {
	passwordHash, err := bcrypthelper.PasswordHash(user.Password)
	if err != nil {
		return &Domain{}, err
	}
	user.Password = passwordHash
	result, err := service.repository.Create(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, err
}

func (service serviceUsers) Login(username string, password string) (string, error) {
	result, err := service.repository.FindByUsername(username)
	if err != nil {
		return "Err", err
	}
	if result.ID == 0 {
		return "Err", err
	}
	if !bcrypthelper.ValidateHash(password, result.Password) {
		return "Err", err
	}
	token := service.jwtAuth.GenerateToken(result.ID)
	return token, err
}

func (service serviceUsers) Edit(user *Domain) (*Domain, error) {
	passwordHash, err := bcrypthelper.PasswordHash(user.Password)
	if err != nil {
		return &Domain{}, err
	}
	user.Password = passwordHash
	result, err := service.repository.Update(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, err
}
