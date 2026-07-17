package core

type IncorrectPasswordErr struct{}

func (e *IncorrectPasswordErr) Error() string {
	return "Incorrect password"
}
