package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSkillRank struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    int
	User      User
	SkillID   int
	Skill     Skill
	Rank      uint `json:"rank"`
}

func NewUserSkillRank() *UserSkillRank {
	return &UserSkillRank{}
}
