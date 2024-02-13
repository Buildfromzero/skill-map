package handlers

import (
	"net/http"

	"github.com/buildfromzero/skill-map/managers"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateUser(ctx *gin.Context)
}

type UserHandler struct {
	group   string
	manager *managers.UserManager
}

func NewUserHandler(manager *managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/",
		manager,
	}
}

func (userH *UserHandler) Register(r *gin.Engine) {
	group := r.Group(userH.group)
	group.GET("users", userH.ListAll)
}

func (userH *UserHandler) ListAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
