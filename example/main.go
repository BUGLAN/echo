package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	engine := echo.New()
	engine.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello world")
	})

	log.Fatal(engine.Start(":8080"))
}
