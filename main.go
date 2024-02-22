package main

import (
	"fmt"

	"github.com/buildfromzero/skill-map/handlers"
	"github.com/buildfromzero/skill-map/managers"
	"github.com/buildfromzero/skill-map/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	storage.InitializeDatabase()
}

func main() {
	fmt.Println("Stating Skill Map...")

	router := gin.Default()
	router.Use(cors.Default())

	userManger := managers.NewUserManager()
	skillManger := managers.NewSkillManager()

	userHandler := handlers.NewUserHandlerFrom(userManger)
	userHandler.RegisterEndpoints(router)

	skillHandler := handlers.NewSkillHandlerFrom(skillManger)
	skillHandler.RegisterEndpoints(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
