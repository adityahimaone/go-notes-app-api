package users

import (
	"gorm.io/gorm"
	"todolist/bussiness/users"
)

type Users struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Name     string
	DOB      string
}

func toDomain(record Users) users.Domain {
	return users.Domain{
		ID:       int(record.ID),
		Username: record.Username,
		Password: record.Password,
		Name:     record.Name,
		DOB:      record.DOB,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		ID:       uint(domain.ID),
		Username: domain.Username,
		Password: domain.Password,
		Name:     domain.Name,
		DOB:      domain.DOB,
	}
}
