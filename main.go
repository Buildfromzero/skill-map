package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Skill Map")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "skill-map-home-page",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
