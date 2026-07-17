package user

type UserRegistrationData struct {
	Username  string
	Email     string
	Password  string
	FirstName *string
	LastName  *string
}

type UserProfile struct {
	Username  string
	FirstName *string
	LastName  *string
	Bio       string
	Country   *string
	City      *string
}
