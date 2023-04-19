package httprest

import "github.com/gofiber/fiber/v2"

type App struct {
	FiberApp *fiber.App
}

func (a App) Routes() {
	a.FiberApp.Post("/decks", func(c *fiber.Ctx) error {
		return a.Create(c)
	})

	a.FiberApp.Get("/decks/:id", func(c *fiber.Ctx) error {
		return a.Create(c)
	})

	a.FiberApp.Put("/decks/:id/draw", func(c *fiber.Ctx) error {
		return a.Draw(c)
	})
}
