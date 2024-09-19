package errors

type EntityInUseError struct {
	message string
}

func NewEntityInUseError(message string) *EntityInUseError {
	return &EntityInUseError{message: message}
}

func (e *EntityInUseError) Error() string {
	return e.message
}
