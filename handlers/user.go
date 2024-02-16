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
	userGroup.DELETE(":userid/", userHandler.Delete)
	userGroup.PATCH(":userid/", userHandler.Update)
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {

	userData := common.NewUserCreationInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		common.BadResponse(ctx, "Failed to bind data")
		return
	}

	newUser, err := userHandler.userManager.Create(userData)

	if err != nil {
		common.BadResponse(ctx, "failed to create user")
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}

func (userHandler *UserHandler) List(ctx *gin.Context) {

	allUsers, err := userHandler.userManager.List()

	if err != nil {
		common.BadResponse(ctx, "failed to get users")
		return
	}

	ctx.JSON(http.StatusOK, allUsers)
}

func (userHandler *UserHandler) Detail(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		fmt.Println("invalid userid")
	}
	user, err := userHandler.userManager.Get(userId)

	if user.ID == 0 {

		common.BadResponse(ctx, "no user present")
		return
	}

	if err != nil {
		fmt.Println("failed to get user")
	}

	ctx.JSON(http.StatusOK, user)
}

func (userHandler *UserHandler) Delete(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		fmt.Println("invalid userid")
	}
	err := userHandler.userManager.Delete(userId)

	if err != nil {
		fmt.Println("failed to get user")
	}

	common.SuccessResponse(ctx, "Deleted user")
}

func (userHandler *UserHandler) Update(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		fmt.Println("invalid userid")
	}

	userData := common.NewUserUpdateInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		common.BadResponse(ctx, "Failed to bind data")
		return
	}

	user, err := userHandler.userManager.Update(userId, userData)

	if err != nil {
		common.BadResponse(ctx, "failed to update user")
		return
	}

	ctx.JSON(http.StatusOK, user)
}
