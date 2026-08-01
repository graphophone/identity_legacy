package postgresdb

type NotFoundErr struct{}

func (e *NotFoundErr) Error() string {
	return "Record not found"
}
