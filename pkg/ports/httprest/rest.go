package httprest

import (
	"github.com/Jonss/cartaman/pkg/usecase/decks"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	FiberApp    *fiber.App
	DeckService decks.DeckService
}

func (a App) Routes() {
	a.FiberApp.Post("/decks", func(c *fiber.Ctx) error {
		return a.Create(c)
	})

	a.FiberApp.Get("/decks/:id", func(c *fiber.Ctx) error {
		return a.Open(c)
	})

	a.FiberApp.Patch("/decks/:id/draw/:count", func(c *fiber.Ctx) error {
		return a.Draw(c)
	})
}
