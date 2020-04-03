package main

import (
	"log"

	"github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/adapters"
	"github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/applications/usecase"
	"github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/drivers"
	"github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/pkg/db"
)

var (
	_mongoDriver       = db.NewMongoDriver()
	_userRepository    = adapters.NewUserGateway(_mongoDriver)
	_userCreateUsecase = usecase.NewUserCreateInteractor(_userRepository)
	_userGetUsecase    = usecase.NewUserGetInteractor(_userRepository)
	_userUpdateUsecase = usecase.NewUserUpdateInteractor(_userRepository)
	_userController    = adapters.NewUserController(_userCreateUsecase, _userGetUsecase, _userUpdateUsecase)
	_server            = drivers.NewServerDriver(_userController)
)

func main() {
	if err := _server.Run(); err != nil {
		log.Fatal(err)
	}
}
