package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

func main() {
	fmt.Println("Skill Map")

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("api/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login-required-testing",
		})
	})

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "skill-map-home-page",
		})
	})

	r.GET("/api/users", func(c *gin.Context) {

		user := User{
			ID:       1,
			FullName: "Tom Victor",
			Email:    "tom@builfromzero.com",
		}

		users := []User{user}
		c.JSON(200, users)
	})

	r.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	r.Run() // listen and serve on 0.0.0.0:8080
}
