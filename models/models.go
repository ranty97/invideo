package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId       string `gorm:"unique"`
	Email        string `gorm:"unique"`
	AvatarURL    string
	AccessToken  string
	RefreshToken string
}
