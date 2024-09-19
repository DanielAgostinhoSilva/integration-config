package errors

type BusinessError struct {
	message string
}

func NewBusinessError(message string) *BusinessError {
	return &BusinessError{message: message}
}

func (e *BusinessError) Error() string {
	return e.message
}
