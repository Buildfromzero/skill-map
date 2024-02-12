package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	route := r.Group("users")
	route.GET("/api/users", createUser)
}
