package main

import (
	"fmt"

	"github.com/buildfromzero/skill-map/database"
	"github.com/buildfromzero/skill-map/handlers"
	"github.com/buildfromzero/skill-map/managers"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Initialize()
}

func main() {
	fmt.Println("Skill Map")

	router := gin.Default()

	userManger := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManger)
	userHandler.RegisterUserApis(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
