package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func UserInfo(c echo.Context) error {
	req := map[string]interface{}{
		"status": "OK",
		"user": map[string]interface{}{
			"token": "TOKEN",
		},
	}

	return c.JSON(http.StatusOK, req)
}

func UserAuth(c echo.Context) error {
	req := map[string]interface{}{
		"status": "OK",
		"token":  "TOKEN",
	}

	return c.JSON(http.StatusCreated, req)
}

func UserLogout(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
