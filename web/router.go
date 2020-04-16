package web

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/skar404/turnip-exchange/handlers"
	"github.com/skar404/turnip-exchange/utils"
)

type Config struct {
	Debug bool
	Host  string
	Port  int
}

func Init() (e *Config) {
	e = &Config{
		Debug: false,
		Host:  "127.0.0.1",
		Port:  8080,
	}

	return
}

func (c *Config) Run() {
	e := echo.New()
	e.Debug = c.Debug

	telegramHook := fmt.Sprintf("/telegram/hook/%s", utils.RandStringRunes(10))

	e.Logger.Printf("telegram hook=%s", telegramHook)
	e.POST(telegramHook, handlers.WebHook)

	api := e.Group("/api/")

	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.GET("ws", handlers.EventWebSocket)

	user := api.Group("user/")
	user.GET("me", handlers.UserInfo)
	user.POST("auth", handlers.UserAuth)
	user.POST("logout", handlers.UserLogout)

	room := api.Group("room/")
	room.GET("room", handlers.GetRooms)
	room.GET("room/:id", handlers.GetRoomById)
	room.POST("room", handlers.CreateRoom)
	room.PATCH("room/:id", handlers.UpdateRoom)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%v", c.Host, c.Port)))
}
