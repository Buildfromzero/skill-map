package models

import (
	"time"

	"gorm.io/gorm"
)

type Competence struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    int            `json:"-"`
	User      User           `json:"-"`
	SkillID   int            `json:"-"`
	Skill     Skill          `json:"skill"`
	Rank      int            `json:"rank"`
}

func NewCompetence() *Competence {
	return &Competence{}
}
