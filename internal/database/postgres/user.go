package postgres

type UserModel struct {
	Username string
	Email string
	PasswordHash string
	FirstName string
	LastName string
	Bio string
	Country *string
	City *string
	ProfileUrl *string
}
