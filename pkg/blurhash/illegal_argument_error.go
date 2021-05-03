package blurhash

type IllegalArgumentError struct {
	detailedMessage string
}

func (e *IllegalArgumentError) Error() string {
	return e.detailedMessage
}

func NewIllegalArgumentError(message string) *IllegalArgumentError {
	return &IllegalArgumentError{detailedMessage: message}
}
