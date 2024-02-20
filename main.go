package main

import (
	"fmt"

	"github.com/buildfromzero/skill-map/handlers"
	"github.com/buildfromzero/skill-map/managers"
	"github.com/buildfromzero/skill-map/storage"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	storage.InitializeDatabase()
}

func main() {
	fmt.Println("Skill Map")

	router := gin.Default()

	router.Use(cors.Default())

	userManger := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManger)
	userHandler.RegisterUserApis(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
