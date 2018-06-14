package serv

import (
	"github.com/YMhao/EasyApi/common"
)

// NewError new an error, and the default code is "default"
func NewDefaultError(err error) *APIError {
	return &APIError{
		Code:    common.ERROR_TYPE_DEFAULT,
		Message: err.Error(),
	}
}

func NewError(code string, err error) *APIError {
	return &APIError{
		Code:    code,
		Message: err.Error(),
	}
}
