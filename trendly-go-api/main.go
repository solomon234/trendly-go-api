package main

import (
	"trendly-go-api/app"
	"trendly-go-api/config"
)

func main() {
	//Test Purposes

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
