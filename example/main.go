package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	engine := echo.New()
	engine.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello world")
	})

	log.Fatal(engine.Start(":8080"))
}
