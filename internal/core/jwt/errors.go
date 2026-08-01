package jwtcore

type InvalidJwt struct{}

func (e *InvalidJwt) Error() string {
	return "Invalid token"
}

type ExpiredJwt struct{}

func (e *ExpiredJwt) Error() string {
	return "Token expired"
}
