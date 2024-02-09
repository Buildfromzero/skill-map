package main

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Skill Map")

	r := gin.Default()

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

	r.Use(static.Serve("/", static.LocalFile("./skillmap/build", true)))

	r.Run() // listen and serve on 0.0.0.0:8080
}
