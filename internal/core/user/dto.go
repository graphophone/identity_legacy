package user

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

type UpdateUserData struct {
	Id        uint
	Username  string
	FirstName *string
	LastName  *string
	Bio       string
	Country   *string
	City      *string
}
