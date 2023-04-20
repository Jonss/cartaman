package httprest

import (
	"github.com/Jonss/cartaman/pkg/usecases/decks"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	FiberApp    *fiber.App
	DeckUseCase decks.DeckUseCase
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
