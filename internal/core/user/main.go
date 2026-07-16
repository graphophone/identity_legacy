package user

type UserManager interface {
	Register(regData UserRegistrationData) error
	UpdateProfile(id int, updData UserUpdateData) error
	UpdatePassword(id int, oldPass, newPass string) error
	Deactive(id int) error
}
