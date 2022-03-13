package errors

type ServiceError struct {
	Code    string `json:"code" example:"MY_CODE"`
	Message string `json:"message" example:"some error message"`
}

// Error show error message.
func (e ServiceError) Error() string {
	return e.Message
}

var (
	ErrValueNotFound = &ServiceError{
		Code:    "VALUE_NOT_FOUND",
		Message: "value not found",
	}

	ErrInternalServerError = &ServiceError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "internal server error",
	}
)
