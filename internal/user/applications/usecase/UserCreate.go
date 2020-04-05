package usecase

import "github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"

type (
	// UserCreateRequset is struct
	UserCreateRequset struct {
		name string
	}

	// UserCreateUsecase is interface
	UserCreateUsecase interface {
		Execute(req UserCreateRequset) (token string, err error)
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
func (p UserCreateInteractor) Execute(req UserCreateRequset) (token string, err error) {
	token, err = p.UserRepository.Create(req.name)

	if err != nil {
		return "", err
	}
	return token, nil
}
