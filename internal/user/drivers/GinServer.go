package drivers

import (
	"fmt"
	"os"

	"github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/adapters"
	"github.com/gin-gonic/gin"
)

var (
	_defaultPort = os.Getenv("PORT")
)

type (
	// IServerDriver is interface
	IServerDriver interface {
		Run() error
	}
	// ServerDriver is struct
	ServerDriver struct {
		Router *gin.Engine
		adapters.IUserController
	}
)

// NewServerDriver is func
func NewServerDriver(userController adapters.IUserController) IServerDriver {
	result := &ServerDriver{}
	result.Router = gin.Default()
	result.IUserController = userController
	return result
}

// CreateRouting is func
func (s ServerDriver) CreateRouting() {
	v1 := s.Router.Group("/v1")
	{
		v1.GET("/user", s.IUserController.Get)
		v1.POST("/user", s.IUserController.Create)
		v1.PUT("/user", s.IUserController.Update)
	}
}

// Run is func
func (s ServerDriver) Run() error {
	fmt.Printf("info: listen on #{_defaultPort}")
	return s.Router.Run(_defaultPort)
}
