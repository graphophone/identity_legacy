package context

type InvalidContextType struct{}

func (e *InvalidContextType) Error() string {
	return "Invalid context type"
}
