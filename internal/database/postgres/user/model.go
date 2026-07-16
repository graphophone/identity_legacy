package userdb

import (
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	Bio          string `gorm:"default:''"`
	Country      *string
	City         *string
	AvatarUrl    *string
	IsActive     bool `gorm:"default:true"`
}

func (u *User) String() string {
	out, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return string(out)
}
