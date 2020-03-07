package usecase

import "github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"

type (
	// UserGetRequest is struct
	UserGetRequest struct {
		token string
	}

	// UserGetResponse is struct
	UserGetResponse struct {
		name string
	}

	// UserGetUsecase is interface
	UserGetUsecase interface {
		Execute(req UserGetRequest) (UserGetResponse, error)
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
func (p UserGetInteractor) Execute(req UserGetRequest) (UserGetResponse, error) {
	user, err := p.UserRepository.Get(req.token)

	if err != nil {
		return UserGetResponse{name: ""}, err
	}
	return UserGetResponse{name: user.Name}, nil
}
