package user

type UserRegistrationData struct {
	Username  string
	Email     string
	Password  string
	FirstName *string
	LastName  *string
}

type UserUpdateData struct {
	Username  *string
	FirstName *string
	LastName  *string
	City      *string
	Country   *string
	Bio       *string
}
