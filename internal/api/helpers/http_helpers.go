package helpers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var errorsMap = map[string]int{
	"BAD_REQUEST": http.StatusBadRequest,
}

func BuildErrorResponse(c echo.Context, errMsg string) error {
	for key, statusCode := range errorsMap {
		if strings.Contains(errMsg, key) {
			return c.JSON(statusCode, map[string]any{
				"error": errMsg,
			})
		}
	}

	return c.JSON(http.StatusInternalServerError, map[string]any{
		"error": errMsg,
	})
}
