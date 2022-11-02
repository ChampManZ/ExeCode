package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func HandleErrorFromPiston(c echo.Context, statusCode int, err error) error {
	if statusCode == -1 {
		statusCode = http.StatusInternalServerError
	}
	return c.JSON(statusCode, ErrorResponse{err.Error()})
}
