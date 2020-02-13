package usecase

import "github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"

type (
	UserGetRequest struct {
		token string
	}

	UserGetResponse struct {
		name string
	}

	UserGetUsecase interface {
		Execute(req UserGetRequest) (UserGetResponse, error)
	}

	UserGetInteractor struct {
		domains.UserRepository
	}
)

func NewUserGetInteractor(userRepository domains.UserRepository) UserGetUsecase {
	result := &UserGetInteractor{}
	result.UserRepository = userRepository
	return result
}

func (p UserGetInteractor) Execute(req UserGetRequest) (UserGetResponse, error) {
	user, err := p.UserRepository.Get(req.token)

	if err != nil {
		return UserGetResponse{name: ""}, err
	}
	return UserGetResponse{name: user.Name}, nil
}
