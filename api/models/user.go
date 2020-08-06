package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Token string `gorm:"size:255;not null;unique" json:"token"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Picture string `gorm:"size:255;not null;unique" json:"picture"`
	UserID string `gorm:"size:100;not null;unique" json:"user_id"`
}

type UserData struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Picture string `json:"picture"`
}