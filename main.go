package main

import (
	"fmt"

	"github.com/buildfromzero/skill-map/handlers"
	"github.com/buildfromzero/skill-map/managers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Skill Map")

	router := gin.Default()

	userManager := managers.NewUserManagerFrom()
	userHandler := handlers.NewUserHandler(userManager)
	userHandler.Register(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
