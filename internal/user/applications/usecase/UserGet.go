package usecase

import "github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"

type (
	// UserGetUsecase is interface
	UserGetUsecase interface {
		Execute(token string) (user *domains.User, err error)
	}

	// UserGetInteractor is struct
	UserGetInteractor struct {
		domains.UserRepository
	}
)

// NewUserGetInteractor is func
func NewUserGetInteractor(userRepository domains.UserRepository) UserGetUsecase {
	result := &UserGetInteractor{}
	result.UserRepository = userRepository
	return result
}

// Execute is func
func (p UserGetInteractor) Execute(token string) (user *domains.User, err error) {
	user, err = p.UserRepository.Get(token)

	if err != nil {
		return nil, err
	}
	return user, nil
}
