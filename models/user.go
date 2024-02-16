package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	FullName  string         `json:"fullName"`
	Email     string         `json:"email"`
}
