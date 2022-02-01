package main

import (
	"QuisAPI/config"
	"QuisAPI/route"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()

	e := echo.New()
	route.NewUser(e)
	route.NewQuestion(e)
	route.NewAnswer(e)

	e.Start(":8080")
}
