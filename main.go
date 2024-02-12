package main

import (
	"fmt"

	"github.com/buildfromzero/skill-map/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Skill Map")

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// // Migrate the schema
	// db.AutoMigrate(&user.User{})

	// db.Create(&user.User{FullName: "Tom Victor", Email: "tom@buildfromzero.com"})

	router := gin.Default()
	router.Use(cors.Default())

	user.RegisterUserRoutes(router)

	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))
	router.Run() // listen and serve on 0.0.0.0:8080
}
