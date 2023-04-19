package main

import (
	"github.com/Jonss/cartaman/pkg/ports/httprest"
	"github.com/gofiber/fiber/v2"
)

func main() {
	r := httprest.App{
		FiberApp: fiber.New(),
	}
	r.Routes()
	// TODO: add logs
	// TODO: get port from config
	r.FiberApp.Listen(":8082")
}
