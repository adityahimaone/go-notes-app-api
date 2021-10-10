package users

type serviceUsers struct {
	repository Repository
}

func NewService(repoUser Repository) Service {
	return &serviceUsers{
		repository: repoUser,
	}
}

func (service serviceUsers) Register(user *Domain) (*Domain, error) {
	result, err := service.repository.Create(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, err
}

func (service serviceUsers) Login(username string, password string) (string, error) {
	result, err := service.repository.FindByUsername(username)
	if err != nil {
		return "&Domain{}", err
	}
	return "result", err
}

func (service serviceUsers) Edit(user *Domain) (*Domain, error) {
	result, err := service.repository.Update(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, err
}
