package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func WebHook(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func EventWebSocket(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
				return
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				return
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
