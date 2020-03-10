package usecase

import "github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"

type (
	// UserUpdateRequest is struct
	UserUpdateRequest struct {
		user domains.User
	}
	// UserUpdateUsecase is interface
	UserUpdateUsecase interface {
		Execute(req UserUpdateRequest) error
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
func (u UserUpdateInteractor) Execute(req UserUpdateRequest) error {
	err := u.UserRepository.Update(req.user)
	if err != nil {
		return err
	}
	return nil
}
