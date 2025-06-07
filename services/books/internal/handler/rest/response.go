package rest

import (
	"github.com/labstack/echo/v4"
)

type (
	ResponseMessage map[string]string

	Response struct {
		Code    int             `json:"code"`
		Message ResponseMessage `json:"message"`
		Data    interface{}     `json:"data"`
	}
)

var (
	DefaultMessage = map[string]string{
		"id": "Berhasil",
		"en": "Success",
	}
)

func DefaultResult(ec echo.Context, code int) error {
	return ResultWithData(ec, code, nil)
}

func ResultError(ec echo.Context, code int, err error) error {
	return ec.JSON(code, Response{
		Code: code,
		Message: map[string]string{
			"id": err.Error(),
			"en": err.Error(),
		},
		Data: nil,
	})
}

func ResultWithData(ec echo.Context, code int, data interface{}) error {
	return ec.JSON(code, Response{
		Code:    code,
		Message: DefaultMessage,
		Data:    data,
	})
}
