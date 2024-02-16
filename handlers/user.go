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
	userGroup.GET("", userHandler.List)
	userGroup.GET(":userid/", userHandler.Detail)
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

func (userHandler *UserHandler) List(ctx *gin.Context) {

	allUsers, err := userHandler.userManager.List()

	if err != nil {
		fmt.Println("failed to get users")
	}

	ctx.JSON(http.StatusOK, allUsers)
}

func (userHandler *UserHandler) Detail(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		fmt.Println("invalid userid")
	}
	user, err := userHandler.userManager.Get(userId)

	if err != nil {
		fmt.Println("failed to get user")
	}

	ctx.JSON(http.StatusOK, user)
}
