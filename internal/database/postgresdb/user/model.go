package userdb

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	PasswordHash string
	FirstName    *string
	LastName     *string
	Bio          *string `gorm:"default:''"`
	Country      *string
	City         *string
	AvatarUrl    *string
	IsActive     bool `gorm:"default:true"`
}
