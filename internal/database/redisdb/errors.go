package redisdb

type IncorrectValueFormat struct{}

func (e *IncorrectValueFormat) Error() string {
	return "Incorrect value format"
}

type InvalidValue struct{}

func (e *InvalidValue) Error() string {
	return "Invalid value"
}
