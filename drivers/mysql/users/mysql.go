package users

import (
	"gorm.io/gorm"
	"todolist/bussiness/users"
)

type repositoryUsers struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) users.Repository {
	return &repositoryUsers{
		DB: db,
	}
}

func (repository repositoryUsers) Create(user *users.Domain) (*users.Domain, error) {
	record := fromDomain(*user)
	if err := repository.DB.Create(&record).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(record)
	return &result, nil
}

func (repository repositoryUsers) Update(user *users.Domain, id int) (*users.Domain, error) {
	record := fromDomain(*user)
	if err := repository.DB.Where("id = ?", id).Updates(&record).Error; err != nil {
		return &users.Domain{}, err
	}
	if err := repository.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(record)
	return &result, nil
}

func (repository repositoryUsers) FindByUsername(username string) (*users.Domain, error) {
	record := Users{}
	if err := repository.DB.Where("username = ?", username).First(&record).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(record)
	return &result, nil
}
