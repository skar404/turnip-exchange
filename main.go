package main

import "github.com/skar404/turnip-exchange/web"

func main() {
	app := web.Init()

	app.Debug = true
	app.Run()
}
