package usecase

import "github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"

type (
	// UserCreateUsecase is interface
	UserCreateUsecase interface {
		Execute(name string) (token *string, err error)
	}

	// UserCreateInteractor is struct
	UserCreateInteractor struct {
		domains.UserRepository
	}
)

// NewUserCreateInteractor is func
func NewUserCreateInteractor(userRepository domains.UserRepository) UserCreateUsecase {
	result := &UserCreateInteractor{}
	result.UserRepository = userRepository
	return result
}

// Execute is func
func (p UserCreateInteractor) Execute(name string) (token *string, err error) {
	token, err = p.UserRepository.Create(domains.User{
		ID:   "",
		Name: name,
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
