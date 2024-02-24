package storage

import (
	"github.com/buildfromzero/skill-map/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(
		&models.User{},
		&models.Skill{},
		&models.SkillGroup{},
		&models.Competence{},
	)
}
