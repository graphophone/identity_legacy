package userdb

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	Email        string
	PasswordHash string
	FirstName    *string
	LastName     *string
	Bio          string `gorm:"default:''"`
	Country      *string
	City         *string
	AvatarUrl    *string
	IsActive     bool `gorm:"default:true"`
}
