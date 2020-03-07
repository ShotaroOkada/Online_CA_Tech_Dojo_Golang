package usecase

import "github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"

type (
	// UserCreateRequest is struct
	UserCreateRequest struct {
		name string
	}

	// UserCreateResponse is struct
	UserCreateResponse struct {
		token string
	}

	// UserCreateUsecase is interface
	UserCreateUsecase interface {
		Execute(req UserCreateRequest) (UserCreateResponse, error)
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
func (p UserCreateInteractor) Execute(req UserCreateRequest) (UserCreateResponse, error) {
	token, err := p.UserRepository.Create(domains.User{
		ID:   "",
		Name: req.name,
	})

	if err != nil {
		return UserCreateResponse{token: ""}, err
	}
	return UserCreateResponse{token: token}, nil
}
