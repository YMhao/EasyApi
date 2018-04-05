package serv

const unknown = "unknown"

// NewError new an error, and the default code is "unknown"
func NewError(err error) *APIError {
	return &APIError{
		Code:    unknown,
		Message: err.Error(),
	}
}
