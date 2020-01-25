package usecase

type (
	// UserCreateRequest is struct
	UserCreateRequest struct {
		name string
	}

	UserCreateResponse struct {
		token string
	}

	// UserCreateUsecase is interface
	UserCreateUsecase interface {
		Execute(req UserCreateRequest) (UserCreateResponse, error)
	}

	UserCreateInteractor struct {
		domains
	}
)
