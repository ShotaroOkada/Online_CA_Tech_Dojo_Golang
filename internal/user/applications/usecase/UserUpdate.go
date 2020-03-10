package usecase

import "github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"

type (
	// UserUpdateUsecase is interface
	UserUpdateUsecase interface {
		Execute(user domains.User) error
	}
	// UserUpdateInteractor is struct
	UserUpdateInteractor struct {
		domains.UserRepository
	}
)

// NewUserUpdateInteractor is func
func NewUserUpdateInteractor(userRepo domains.UserRepository) UserUpdateUsecase {
	result := &UserUpdateInteractor{}
	result.UserRepository = userRepo
	return result
}

// Execute is func
func (u UserUpdateInteractor) Execute(user domains.User) error {
	err := u.UserRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}
