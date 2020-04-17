package web

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

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

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetMaxPoolSize(20))
	if err != nil {
		e.Logger.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())

	a := func(c *mongo.Client) {
		err = client.Ping(nil, nil)
		if err != nil {
			fmt.Print(err)
		}
	}

	for true {

		a(client)

		//time.Sleep(1 * time.Second)

		//fmt.Print("start")
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%v", c.Host, c.Port)))
}
