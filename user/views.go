package user

import "github.com/gin-gonic/gin"

func createUser(c *gin.Context) {

	user := UserModel{
		FullName: "Tom Victor",
		Email:    "tom@builfromzero.com",
	}

	users := []UserModel{user}
	c.JSON(200, users)
}
