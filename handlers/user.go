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
	userManager managers.UserManager
}

func NewUserHandlerFrom(userManager managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/users/",
		userManager,
	}
}

func (handler *UserHandler) RegisterEndpoints(r *gin.Engine) {
	userGroup := r.Group(handler.groupName)

	userGroup.GET("", handler.ListUser)
	userGroup.POST("", handler.CreateUser)
	userGroup.POST(":userid/skills", handler.AddSkill)
	userGroup.GET(":userid/", handler.UserDetail)
	userGroup.DELETE(":userid/", handler.DeleteUser)
	userGroup.PATCH(":userid/", handler.UpdateUser)

}

func (handler *UserHandler) CreateUser(ctx *gin.Context) {

	userData := common.NewUserCreationInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		common.BadResponse(ctx, "failed to bind data")
		return
	}

	newUser, err := handler.userManager.Create(userData)

	if err != nil {
		common.BadResponse(ctx, "failed to create user")
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}

func (handler *UserHandler) ListUser(ctx *gin.Context) {

	allUsers, err := handler.userManager.List()

	if err != nil {
		common.BadResponse(ctx, "failed to get users")
		return
	}

	ctx.JSON(http.StatusOK, allUsers)
}

func (handler *UserHandler) UserDetail(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		fmt.Println("invalid userid")
		return
	}
	user, err := handler.userManager.Get(userId)

	if err != nil {
		common.BadResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (handler *UserHandler) DeleteUser(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		common.BadResponse(ctx, "invalid userid")
		return
	}
	err := handler.userManager.Delete(userId)

	if err != nil {
		common.BadResponse(ctx, err.Error())
		return
	}

	common.SuccessResponse(ctx, "Deleted user")
}

func (handler *UserHandler) UpdateUser(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		common.BadResponse(ctx, "failed to delete user")
		return
	}

	userData := common.NewUserUpdateInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		common.BadResponse(ctx, "failed to bind data")
		return
	}

	user, err := handler.userManager.Update(userId, userData)

	if err != nil {
		common.BadResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (handler *UserHandler) AddSkill(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userid")

	if !ok {
		common.BadResponse(ctx, "failed to delete user")
		return
	}

	userData := common.NewCompetenceInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		common.BadResponse(ctx, "failed to bind data")
		return
	}

	user, err := handler.userManager.AddNewSkill(userId, userData)

	if err != nil {
		common.BadResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}
