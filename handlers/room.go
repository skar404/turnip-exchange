package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetRooms(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetRoomById(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func CreateRoom(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func UpdateRoom(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
