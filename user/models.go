package user

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}
