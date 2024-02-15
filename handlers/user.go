package handlers

import (
	"fmt"
	"net/http"

	"github.com/buildfromzero/skill-map/common"
	"github.com/buildfromzero/skill-map/managers"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	userManager *managers.UserManager
}

func NewUserHandlerFrom(userManager *managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/users",
		userManager,
	}
}

func (userHandler *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(userHandler.groupName)
	userGroup.POST("", userHandler.Create)
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {

	userData := common.NewUserCreationInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		fmt.Println("Failed to bind data")
	}

	newUser, err := userHandler.userManager.Create(userData)

	if err != nil {
		fmt.Println("failed to create user")
	}

	ctx.JSON(http.StatusOK, newUser)
}
