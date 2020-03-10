package adapters

import (
	"net/http"

	"github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/applications/usecase"
	"github.com/gin-gonic/gin"
)

type (
	// IUserController is interface
	IUserController interface {
		Create(c *gin.Context)
		Get(c *gin.Context)
		Update(c *gin.Context)
	}
	// UserController is struct
	UserController struct {
		usecase.UserCreateUsecase
		usecase.UserGetUsecase
		usecase.UserUpdateUsecase
	}
)

// NewUserController is func
func NewUserController(create usecase.UserCreateUsecase, get usecase.UserGetUsecase, update usecase.UserUpdateUsecase) IUserController {
	result := &UserController{}
	result.UserCreateUsecase = create
	result.UserGetUsecase = get
	result.UserUpdateUsecase = update
	return result
}

// Create is func
func (u UserController) Create(c *gin.Context) {
	name := c.Param("id")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "parameter name is empty"})
		return
	}
	result, err := u.UserCreateUsecase.Execute(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

// Get is func
func (u UserController) Get(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "parameter token is empty"})
		return
	}
	result, err := u.UserGetUsecase.Execute(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Update is func
func (u UserController) Update(c *gin.Context) {
	var reqObject usecase.UserUpdateRequest
	if err := c.BindJSON(&reqObject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.UserUpdateUsecase.Execute(reqObject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
	return
}
