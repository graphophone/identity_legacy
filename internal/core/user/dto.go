package user

import "encoding/json"

type RegisterUserData struct {
	Username  string
	Email     string
	Password  string
	FirstName *string
	LastName  *string
}

type UserProfile struct {
	Id        uint
	Username  string
	FirstName *string
	LastName  *string
	Bio       string
	Country   *string
	City      *string
	AvatarUrl *string
}

func (u *UserProfile) String() string {
	out, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return string(out)
}

type UpdateProfileData struct {
	Id        uint
	Username  string
	FirstName *string
	LastName  *string
	Bio       string
	Country   *string
	City      *string
}
