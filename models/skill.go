package models

import (
	"time"

	"gorm.io/gorm"
)

func NewSkill() *Skill {
	return &Skill{}
}

func NewSkillGroup() *SkillGroup {
	return &SkillGroup{}
}

type SkillGroup struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
	// Skills    []Skill        `gorm:"many2many:skillgroup_skills;" json:"skills"`
}

type Skill struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"name"`
	SkillGroupID uint           `json:"-"`
	SkillGroup   SkillGroup     `json:"skillGroup"`
}
